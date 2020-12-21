package main

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/jsonschema"
	keptn "github.com/keptn/go-utils/pkg/lib"
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
)

func createDataSection(md *MarkDown, eventType string, data interface{}) {
	md.Title("Data", 4)
	md.WriteLineBreak()
	md.Writeln("<details><summary>Json Schema of " + eventType + "</summary>")
	md.Writeln("<p>")
	md.WriteLineBreak()
	md.CodeBlock(toJsonSchema(data), "json")
	md.Writeln("</p>")
	md.Writeln("</details>")
	md.WriteLineBreak()
}

func createExampleSection(md *MarkDown, eventType string, data interface{}) {
	md.Title("Example", 4)
	md.WriteLineBreak()
	md.CodeBlock(toJson(ce(eventType, data)), "json")
	md.WriteLineBreak()
}
func createUpLink(md *MarkDown) {
	md.UpLink()
}

func createSection(md *MarkDown, title string, eventType string, data interface{}) {
	md.Title(title, 3)
	md.Title("Type", 4)
	md.Writeln(eventType)
	createDataSection(md, eventType, data)
	createExampleSection(md, eventType, data)
	createUpLink(md)
}

func createSectionTitle(md *MarkDown, title string) {
	md.Title(title, 2)
}

func main() {

	md := NewMarkDown()

	md.Title("Keptn Cloud Events", 1)
	md.Bullet().Link("Project", "#project")
	md.Bullet().Link("Service", "#service")
	md.Bullet().Link("Approval", "#approval")
	md.Bullet().Link("Deployment", "#deployment")
	md.Bullet().Link("Test", "#test")
	md.Bullet().Link("Evaluation", "#evaluation")
	md.Bullet().Link("Release", "#release")
	md.Bullet().Link("Remediation", "#remediation")
	md.Bullet().Link("Action", "#action")
	md.Bullet().Link("Get-SLI", "#get-sli")
	md.Bullet().Link("Problem", "#problem")
	md.Writeln("---")
	md.MultiBr(2)

	createSectionTitle(md, "Project")
	createSection(md, "Project Create Triggered", keptnv2.GetTriggeredEventType(keptnv2.ProjectCreateTaskName), projectCreateData)
	createSection(md, "Project Create Started", keptnv2.GetStartedEventType(keptnv2.ProjectCreateTaskName), projectCreateStartedEventData)
	createSection(md, "Project Create Finished", keptnv2.GetFinishedEventType(keptnv2.ProjectCreateTaskName), projectCreateFinishedEventData)

	createSectionTitle(md, "Service")
	createSection(md, "Service Create Started", keptnv2.GetStartedEventType(keptnv2.ServiceCreateTaskName), serviceCreateStartedEventData)
	createSection(md, "Service Create Status Changed", keptnv2.GetStatusChangedEventType(keptnv2.ServiceCreateTaskName), serviceCreateStatusChangesData)
	createSection(md, "Service Create Finished", keptnv2.GetFinishedEventType(keptnv2.ServiceCreateTaskName), serviceCreateFinishedEventData)

	createSectionTitle(md, "Approval")
	createSection(md, "Approval Triggered", keptnv2.GetTriggeredEventType(keptnv2.ApprovalTaskName), approvalTriggeredEventData)
	createSection(md, "Approval Started", keptnv2.GetStartedEventType(keptnv2.ApprovalTaskName), approvalStartedEventData)
	createSection(md, "Approval Status Changed", keptnv2.GetStatusChangedEventType(keptnv2.ApprovalTaskName), approvalStatusChangedEventData)
	createSection(md, "Approval Finished", keptnv2.GetFinishedEventType(keptnv2.ApprovalTaskName), approvalFinishedEventData)

	createSectionTitle(md, "Deployment")
	createSection(md, "Deployment Triggered", keptnv2.GetTriggeredEventType(keptnv2.DeploymentTaskName), deploymentTriggeredEventData)
	createSection(md, "Deployment Started", keptnv2.GetStartedEventType(keptnv2.DeploymentTaskName), deploymentStartedEventData)
	createSection(md, "Deployment Status Changed", keptnv2.GetStatusChangedEventType(keptnv2.DeploymentTaskName), deploymentStatusChangedEventData)
	createSection(md, "Deployment Finished", keptnv2.GetFinishedEventType(keptnv2.DeploymentTaskName), deploymentFinishedEventData)

	createSectionTitle(md, "Test")
	createSection(md, "Test Triggered", keptnv2.GetTriggeredEventType(keptnv2.TestTaskName), testTriggeredEventData)
	createSection(md, "Test Started", keptnv2.GetStartedEventType(keptnv2.TestTaskName), testStartedEventData)
	createSection(md, "Test Status Changed", keptnv2.GetTriggeredEventType(keptnv2.TestTaskName), testStatusChangedEventData)
	createSection(md, "Test Finished", keptnv2.GetFinishedEventType(keptnv2.TestTaskName), testTestFinishedEventData)

	createSectionTitle(md, "Evaluation")
	createSection(md, "Evaluation Triggered", keptnv2.GetTriggeredEventType(keptnv2.EvaluationTaskName), evaluationTriggeredEventData)
	createSection(md, "Evaluation Started", keptnv2.GetStartedEventType(keptnv2.EvaluationTaskName), evaluationStartedEventData)
	createSection(md, "Evaluation Status Changed", keptnv2.GetStatusChangedEventType(keptnv2.EvaluationTaskName), evaluationStatusChangedEventData)
	createSection(md, "Evaluation Finished", keptnv2.GetFinishedEventType(keptnv2.EvaluationTaskName), evaluationFinishedEventData)

	createSectionTitle(md, "Release")
	createSection(md, "Release Triggered", keptnv2.GetTriggeredEventType(keptnv2.ReleaseTaskName), releaseTriggeredEventData)
	createSection(md, "Release Started", keptnv2.GetStartedEventType(keptnv2.ReleaseTaskName), releaseStartedEventData)
	createSection(md, "Release Status Changed", keptnv2.GetStatusChangedEventType(keptnv2.ReleaseTaskName), releaseStatusChangedEventData)
	createSection(md, "Release Finished", keptnv2.GetFinishedEventType(keptnv2.ReleaseTaskName), releaseFinishedEventData)

	createSectionTitle(md, "Remediation")
	remediationTaskName := "remediation" // TODO: define task name in go-utils
	createSection(md, "Remediation Triggered", keptnv2.GetTriggeredEventType(remediationTaskName), remediationTriggeredEventData)
	createSection(md, "Remediation Started", keptnv2.GetStartedEventType(remediationTaskName), remediationStartedEventData)
	createSection(md, "Remediation Status Changed", keptnv2.GetStatusChangedEventType(remediationTaskName), remediationStatusChangedEventData)
	createSection(md, "Remediation Finished", keptnv2.GetFinishedEventType(remediationTaskName), remediationFinishedEventData)

	createSectionTitle(md, "Action")
	createSection(md, "Action Triggered", keptnv2.GetTriggeredEventType(keptnv2.ActionTaskName), actionTriggeredEventData)
	createSection(md, "Action Started", keptnv2.GetStartedEventType(keptnv2.ActionTaskName), actionStartedEventData)
	createSection(md, "Action Finished", keptnv2.GetFinishedEventType(keptnv2.ActionTaskName), actionFinishedEventData)

	createSectionTitle(md, "Get SLI")
	createSection(md, "Get SLI Triggered", keptnv2.GetTriggeredEventType(keptnv2.GetSLITaskName), getSLITriggeredEventData)
	createSection(md, "Get SLI Started", keptnv2.GetStartedEventType(keptnv2.GetSLITaskName), getSLIStartedEventData)
	createSection(md, "Get SLI Finished", keptnv2.GetFinishedEventType(keptnv2.GetSLITaskName), getSLIFinishedEventData)

	createSectionTitle(md, "Monitoring")
	createSection(md, "Configure Monitoring Triggered", keptnv2.GetTriggeredEventType(keptnv2.ConfigureMonitoringTaskName), configureMonitoringTriggeredEventData)
	createSection(md, "Configure Monitoring Started", keptnv2.GetStartedEventType(keptnv2.ConfigureMonitoringTaskName), configureMonitoringStartedEventData)
	createSection(md, "Configure Monitoring Finished", keptnv2.GetFinishedEventType(keptnv2.ConfigureMonitoringTaskName), configureMonitoringFinishedEventData)

	createSectionTitle(md, "Problem")
	createSection(md, "Problem Open", keptn.ProblemOpenEventType, problemOpenEventData)
	fmt.Println(md.String())

}

func toJsonSchema(j interface{}) string {
	schema := jsonschema.Reflect(j)
	schemaStr, _ := json.MarshalIndent(schema, "", "  ")
	return string(schemaStr)

}

func toJson(j interface{}) string {
	jsonStr, _ := json.MarshalIndent(j, "", "  ")
	return string(jsonStr)
}
