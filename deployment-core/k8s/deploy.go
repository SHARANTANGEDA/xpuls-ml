package k8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/deployment-core/types"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
)

// DeploymentService describes the behavior of a deployment processor.
type DeploymentService interface {
	ProcessDeployment(yamlContent []byte, options DeploymentOptions) error
}

// DeploymentOptions contains various deployment parameters.
type DeploymentOptions struct {
	DryRun    bool   `json:"dry_run"`
	Canary    bool   `json:"canary"`
	GithubUrl string `json:"github_url"`
	CommitId  string `json:"commit_id"`
	Traffic   int    `json:"traffic"` // Percentage of traffic to be sent to the canary. Only relevant if Canary is true.
}

type DeployManager struct {
	configuration *types.DeploymentConfiguration
}

func (s *DeployManager) ProcessDeployment(c *gin.Context) {
	var options DeploymentOptions

	if err := c.ShouldBindJSON(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(options.CommitId)

	// For now, just unmarshal to a generic map to ensure the YAML is valid.
	s.configuration, _ = ReadDeploymentConfigFromFile(
		"/Users/sharan/xpuls/xpuls-ml/server-config.yaml")
	fmt.Println(s.configuration.Server.Name)

	// If dryRun, just print the content without doing anything.
	if options.DryRun {
		fmt.Println("Dry Run - Parsed YAML:", s.configuration)
	}

	// Handle canary deployments.
	if options.Canary {
		fmt.Printf("Canary Deployment with %d%% traffic.\n", options.Traffic)
		// Actual logic for canary deployment would go here.
	} else {
		fmt.Println("Standard Deployment.")
		// Actual logic for standard deployment would go here.
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deployment processed successfully!"})
}

func ReadDeploymentConfigFromFile(filename string) (*types.DeploymentConfiguration, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var config types.DeploymentConfiguration
	if err := yaml.Unmarshal(content, &config); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &config, nil
}
