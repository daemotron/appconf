// Package appconf is a lightweight configuration solution for Go applications.
// It helps manage configuration settings and handle configuration inputs from
// different sources.
//
// Currently, this package supports the following configuration sources:
//
//   - JSON Files
//   - Environment Variables
//   - Command Line Flags
//
// Configuration directives are interpreted following this precedence order:
//
//  1. Command Line Flags
//  2. Environment Variables
//  3. Configuration File
//  4. Default Values
//
// Settings provided by an instance with a lower precedence order (i.e. higher priority)
// will always override those with a higher precedence order (i.e. lower priority).
package appconf
