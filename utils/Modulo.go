package utils

// mod returns the result of a modulo b, but always positive
func mod(a, b int) int {
    return (a % b + b) % b
}