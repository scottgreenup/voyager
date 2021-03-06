package ups

import (
	"encoding/json"
	"strings"

	"github.com/atlassian/voyager/pkg/orchestration/wiring/wiringutil"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/wiringutil/osb"
	"k8s.io/apimachinery/pkg/runtime"

	smith_v1 "github.com/atlassian/smith/pkg/apis/smith/v1"
	"github.com/atlassian/voyager"
	orch_v1 "github.com/atlassian/voyager/pkg/apis/orchestration/v1"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/wiringplugin"
	"github.com/atlassian/voyager/pkg/orchestration/wiring/wiringutil/knownshapes"
	sc_v1b1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/pkg/errors"
)

const (
	ResourceType   voyager.ResourceType = "UPS"
	ResourcePrefix                      = "UPS"

	clusterServiceClassExternalID = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468"
	clusterServicePlanExternalID  = "86064792-7ea2-467b-af93-ac9694d96d52"
)

var (
	envVarReplacer    = strings.NewReplacer("-", "_", ".", "_")
	defaultUpsEnvVars = map[string]string{
		"SPECIAL_KEY_1": "data.special-key-1",
		"SPECIAL_KEY_2": "data.special-key-2",
	}
)

type WiringPlugin struct {
}

func New() *WiringPlugin {
	return &WiringPlugin{}
}

func (p *WiringPlugin) WireUp(resource *orch_v1.StateResource, context *wiringplugin.WiringContext) wiringplugin.WiringResult {
	if resource.Type != ResourceType {
		return &wiringplugin.WiringResultFailure{
			Error: errors.Errorf("invalid resource type: %q", resource.Type),
		}
	}

	serviceInstance, err := osb.ConstructServiceInstance(resource, clusterServiceClassExternalID, clusterServicePlanExternalID)
	if err != nil {
		return &wiringplugin.WiringResultFailure{
			Error: err,
		}
	}

	instanceParameters, external, retriable, err := instanceParameters(resource)
	if err != nil {
		return &wiringplugin.WiringResultFailure{
			Error:            err,
			IsExternalError:  external,
			IsRetriableError: retriable,
		}
	}
	if instanceParameters != nil {
		serviceInstance.Spec.Parameters = &runtime.RawExtension{
			Raw: instanceParameters,
		}
	}

	instanceResourceName := wiringutil.ServiceInstanceResourceName(resource.Name)

	smithResource := smith_v1.Resource{
		Name:       instanceResourceName,
		References: nil, // No references
		Spec: smith_v1.ResourceSpec{
			Object: serviceInstance,
		},
	}

	shapes, external, retriable, err := instanceShapes(&smithResource)
	if err != nil {
		return &wiringplugin.WiringResultFailure{
			Error:            err,
			IsExternalError:  external,
			IsRetriableError: retriable,
		}
	}

	return &wiringplugin.WiringResultSuccess{
		Contract: wiringplugin.ResourceContract{
			Shapes: shapes,
		},
		Resources: []smith_v1.Resource{smithResource},
	}
}

func instanceShapes(smithResource *smith_v1.Resource) ([]wiringplugin.Shape, bool /* external */, bool /* retriable */, error) {
	// UPS outputs all of its inputs
	si := smithResource.Spec.Object.(*sc_v1b1.ServiceInstance)
	parameters := map[string]json.RawMessage{}
	if si.Spec.Parameters == nil {
		return []wiringplugin.Shape{
			knownshapes.NewBindableEnvironmentVariables(smithResource.Name, ResourcePrefix, defaultUpsEnvVars),
		}, false, false, nil
	}

	err := json.Unmarshal(si.Spec.Parameters.Raw, &parameters)
	if err != nil {
		return nil, false, false, errors.WithStack(err)
	}

	credentials, ok := parameters["credentials"]
	if !ok {
		return []wiringplugin.Shape{
			knownshapes.NewBindableEnvironmentVariables(smithResource.Name, ResourcePrefix, defaultUpsEnvVars),
		}, false, false, nil
	}

	credentialsMap := map[string]string{}
	err = json.Unmarshal(credentials, &credentialsMap)
	if err != nil {
		return nil, false, false, errors.WithStack(err)
	}

	envVars := map[string]string{}
	for k := range credentialsMap {
		envVars[makeEnvVarName(k)] = "data." + k
	}
	return []wiringplugin.Shape{
		knownshapes.NewBindableEnvironmentVariables(smithResource.Name, ResourcePrefix, envVars),
	}, false, false, nil
}

func makeEnvVarName(elements ...string) string {
	return strings.ToUpper(envVarReplacer.Replace(strings.Join(elements, "_")))
}

// Just a straight passthrough...
// (should probably just implement a default autowiring function similar to how RPS OSB works, which
// takes the class/plan names as arguments?)
func instanceParameters(resource *orch_v1.StateResource) ([]byte, bool, bool, error) {
	if resource.Spec == nil {
		return nil, false, false, nil
	}

	return resource.Spec.Raw, false, false, nil
}
