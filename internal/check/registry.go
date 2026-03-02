package check

import "github.com/vidya381/devcheck/internal/detector"

func Build(stack detector.DetectedStack) []Check {
	var cs []Check

	if stack.Go {
		cs = append(cs, &BinaryCheck{Binary: "go"})
		cs = append(cs, &GoVersionCheck{Dir: "."})
	}
	if stack.Node {
		cs = append(cs, &BinaryCheck{Binary: "node"})
		cs = append(cs, &BinaryCheck{Binary: "npm"})
		cs = append(cs, &NodeVersionCheck{Dir: "."})
	}
	if stack.Python {
		cs = append(cs, &BinaryCheck{Binary: "python3"})
		cs = append(cs, &BinaryCheck{Binary: "pip"})
	}
	if stack.Java {
		cs = append(cs, &BinaryCheck{Binary: "java"})
		if stack.Maven {
			cs = append(cs, &BinaryCheck{Binary: "mvn"})
		}
		if stack.Gradle {
			cs = append(cs, &BinaryCheck{Binary: "gradle"})
		}
	}
	if stack.Docker {
		cs = append(cs, &BinaryCheck{Binary: "docker"})
		cs = append(cs, &DockerDaemonCheck{})
	}
	if stack.Postgres {
		// add Postgres checks
	}
	if stack.Redis {
		// add Redis checks
	}

	if stack.EnvExample {
		cs = append(cs, &EnvCheck{Dir: "."})
	}

	return cs
}
