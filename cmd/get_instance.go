package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

func init() {
	var compartmentId, displayName, lifeCycleState, availabilityDomain string

	var getInstanceCmd = &cobra.Command{
		Use:   "instance",
		Short: "TODO",
		Long:  "TODO",

		Run: func(cmd *cobra.Command, args []string) {

			listOfInstances := getInstance(compartmentId, displayName, lifeCycleState, availabilityDomain)

			fmt.Printf("%-50s%-20s%-20s%-80s\n", "Display Name", "Lifecycle State", "Shape", "OCID")
			for _, item := range listOfInstances {
				fmt.Printf("%-50s%-20s%-20s%-80s\n", *item.DisplayName, item.LifecycleState, *item.Shape, *item.Id)
			}

		},
	}
	getInstanceCmd.Flags().StringVarP(&displayName, "name", "n", "", "name of the vcn")
	getInstanceCmd.Flags().StringVarP(&lifeCycleState, "lifecycle-state", "l", "", "lifecycle state")
	getInstanceCmd.Flags().StringVarP(&compartmentId, "compartment-id", "i", os.Getenv("TF_VAR_compartment_ocid"), "ocid of compartment")
	getInstanceCmd.Flags().StringVarP(&availabilityDomain, "availability-domain", "a", "", "availability domain")

	GetCmd.AddCommand(getInstanceCmd)
}

func getInstance(compartmentId, displayName, lifeCycleState, availabilityDomain string) []core.Instance {
	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Print("Error: ")
		panic(err)
	}

	ctx := context.Background()

	request := core.ListInstancesRequest{}

	request.CompartmentId = &compartmentId

	if displayName != "" {
		request.DisplayName = &displayName
	}

	if availabilityDomain != "" {
		request.AvailabilityDomain = &availabilityDomain
	}

	if strings.ToLower(lifeCycleState) == "running" {
		request.LifecycleState = core.InstanceLifecycleStateRunning
	} else if strings.ToLower(lifeCycleState) == "provisioning" {
		request.LifecycleState = core.InstanceLifecycleStateProvisioning
	} else if strings.ToLower(lifeCycleState) == "stopped" {
		request.LifecycleState = core.InstanceLifecycleStateStopped
	} else if strings.ToLower(lifeCycleState) == "terminated" {
		request.LifecycleState = core.InstanceLifecycleStateTerminated
	} else if strings.ToLower(lifeCycleState) == "starting" {
		request.LifecycleState = core.InstanceLifecycleStateStarting
	} else if strings.ToLower(lifeCycleState) == "stopping" {
		request.LifecycleState = core.InstanceLifecycleStateStopping
	} else if strings.ToLower(lifeCycleState) == "terminating" {
		request.LifecycleState = core.InstanceLifecycleStateTerminating
	}

	response, err := client.ListInstances(ctx, request)
	if err != nil {
		fmt.Println("error at get instances")
		panic(err)
	}

	return response.Items
}