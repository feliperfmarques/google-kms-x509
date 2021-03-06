package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"

	"github.com/spf13/cobra"
)

var emailAddressOID = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}

var (
	commonName         string
	country            string
	province           string
	locality           string
	organization       string
	organizationalUnit string
	emailAddress       string
)

func addSubjectFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&commonName, "common-name", "", "x509 Distinguished Name (DN) field",
	)

	cmd.Flags().StringVar(
		&country, "country", "", "x509 Distinguished Name (DN) field",
	)

	cmd.Flags().StringVar(
		&province, "province", "", "x509 Distinguished Name (DN) field",
	)

	cmd.Flags().StringVar(
		&locality, "locality", "", "x509 Distinguished Name (DN) field",
	)

	cmd.Flags().StringVar(
		&organization, "organization", "", "x509 Distinguished Name (DN) field",
	)

	cmd.Flags().StringVar(
		&organizationalUnit, "organizationalUnit", "", "x509 Distinguished Name (DN) field",
	)

	cmd.Flags().StringVar(
		&emailAddress, "emailAddress", "", "x509 Distinguished Name (DN) field",
	)

	cmd.MarkFlagRequired("common-name")
}

func convertSubjectFlagsToName() pkix.Name {
	name := pkix.Name{
		CommonName: commonName,
	}

	if country != "" {
		name.Country = []string{country}
	}

	if province != "" {
		name.Province = []string{province}
	}

	if locality != "" {
		name.Locality = []string{locality}
	}

	if organization != "" {
		name.Organization = []string{organization}
	}

	if organizationalUnit != "" {
		name.OrganizationalUnit = []string{organizationalUnit}
	}

	if emailAddress != "" {
		name.ExtraNames = []pkix.AttributeTypeAndValue{
			{
				Type:  emailAddressOID,
				Value: emailAddress,
			},
		}
	}

	return name
}
