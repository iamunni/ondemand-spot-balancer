package scheduler

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const SchedulerName = "OndemandSpotBalancer"

// NetworkOverhead : Filter and Score nodes based on Pod's AppGroup requirements: MaxNetworkCosts requirements among Pods with dependencies
type OndemandSpotBalancer struct {
	client.Client
	podLister corelisters.PodLister
	handle    framework.Handle
}

// PreFilterState computed at PreFilter and used at Filter and Score.
type PreFilterState struct {
	//
	ondemandPerentage int64
	spotPerentage     int64

	//
	calculatedPercentage map[string]map[string]int64
}

var _ framework.PreFilterPlugin = &OndemandSpotBalancer{}
var _ framework.FilterPlugin = &OndemandSpotBalancer{}
var _ framework.ScorePlugin = &OndemandSpotBalancer{}

func (*OndemandSpotBalancer) Name() string {
	return SchedulerName
}

// NewScheduler initializes a new plugin and returns it.
func NewScheduler(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &OndemandSpotBalancer{}, nil
}

// PreFilter performs the following operations:
// 1. Get the owner
// 2. Get desired % of ondemand/spot from deployment annotation
// 3. Get the list of pods belonging to the app
// 4. Get the list of nodes running the pods
// 5. Calculate the current % in ondemand/spot
// 6. Calculate the % in ondemand/spot, after adding new po
// 7. Save desired % and calculated % to PreFilterState
func (osb *OndemandSpotBalancer) PreFilter(ctx context.Context, state *framework.CycleState, pod *corev1.Pod) (*framework.PreFilterResult, *framework.Status) {
	return nil, nil
}

// PreFilterExtensions is used for updating plugin state when a pod is added or removed (for preemption).
// This plugin is stateless, so these functions are not needed
func (osb *OndemandSpotBalancer) PreFilterExtensions() framework.PreFilterExtensions { return s }

func (*OndemandSpotBalancer) AddPod(ctx context.Context, cycleState *framework.CycleState, podToSchedule *corev1.Pod, podToAdd *framework.PodInfo, nodeInfo *framework.NodeInfo) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

func (*OndemandSpotBalancer) RemovePod(ctx context.Context, cycleState *framework.CycleState, podToSchedule *corev1.Pod, podToRemove *framework.PodInfo, nodeInfo *framework.NodeInfo) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

// Filter performs the following operations:
// 1. Get the % in ondemand/spot from PreFilterState
// 2. Get the node type (ondemand/spot)
// 2. Check if the pod can be added to the node. If ondemand/spot % is met, allow. otherwise fail
func (osb *OndemandSpotBalancer) Filter(ctx context.Context,
	cycleState *framework.CycleState,
	pod *corev1.Pod,
	nodeInfo *framework.NodeInfo) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

// Score performs the following operations:
// 1. Get the % in ondemand/spot from PreFilterState
// 2. Get the type of current node
// 3. Set the score for the node. score = calcalated % - desried %
func (osb *OndemandSpotBalancer) Score(ctx context.Context,
	cycleState *framework.CycleState,
	pod *corev1.Pod,
	nodeName string) (int64, *framework.Status) {
	return 0, framework.NewStatus(framework.Success, "")
}

// NormalizeScore : normalize scores since lower scores correspond to lower priority
func (osb *OndemandSpotBalancer) NormalizeScore(ctx context.Context,
	state *framework.CycleState,
	pod *corev1.Pod,
	scores framework.NodeScoreList) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

// ScoreExtensions : an interface for Score extended functionality
func (osb *OndemandSpotBalancer) ScoreExtensions() framework.ScoreExtensions {
	return osb
}
