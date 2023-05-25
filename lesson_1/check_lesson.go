package test

import (
    "testing"
    "os/exec"
    "strings"
    "github.com/stretchr/testify/assert"
)

func TestProvisioners(t *testing.T) {
    // terraform init and apply
    terraformInit := exec.Command("terraform", "init")
    terraformApply := exec.Command("terraform", "apply", "-auto-approve")
    terraformDestroy := exec.Command("terraform", "destroy", "-force")

    errInit := terraformInit.Run()
    if errInit != nil {
        t.Error(errInit)
    }

    // capture stdout/stderr of terraform apply
    var stdout, stderr strings.Builder
    terraformApply.Stdout = &stdout
    terraformApply.Stderr = &stderr

    errApply := terraformApply.Run()
    if errApply != nil {
        t.Error(errApply)
    }

    // check stdout for "Hello, World!"
    assert.Contains(t, stdout.String(), "Hello, World!")

    // terraform destroy
    errDestroy := terraformDestroy.Run()
    if errDestroy != nil {
        t.Error(errDestroy)
    }
}
