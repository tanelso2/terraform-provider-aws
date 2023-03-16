package rds_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/rds"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccRDSClusterDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_rds_cluster.test"
	resourceName := "aws_rds_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "backtrack_window", resourceName, "backtrack_window"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cluster_identifier", resourceName, "cluster_identifier"),
					resource.TestCheckResourceAttrPair(dataSourceName, "database_name", resourceName, "database_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "db_cluster_parameter_group_name", resourceName, "db_cluster_parameter_group_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "db_subnet_group_name", resourceName, "db_subnet_group_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine", resourceName, "engine"),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine_mode", resourceName, "engine_mode"),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine_version", resourceName, "engine_version"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hosted_zone_id", resourceName, "hosted_zone_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "master_username", resourceName, "master_username"),
					resource.TestCheckResourceAttrPair(dataSourceName, "network_type", resourceName, "network_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags.Name", resourceName, "tags.Name"),
				),
			},
		},
	})
}

func TestAccRDSClusterDataSource_manage_password(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_rds_cluster.test"
	resourceName := "aws_rds_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterDataSourceConfig_manage_password(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "backtrack_window", resourceName, "backtrack_window"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cluster_identifier", resourceName, "cluster_identifier"),
					resource.TestCheckResourceAttrPair(dataSourceName, "database_name", resourceName, "database_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "db_cluster_parameter_group_name", resourceName, "db_cluster_parameter_group_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "db_subnet_group_name", resourceName, "db_subnet_group_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine", resourceName, "engine"),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine_mode", resourceName, "engine_mode"),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine_version", resourceName, "engine_version"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hosted_zone_id", resourceName, "hosted_zone_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "master_username", resourceName, "master_username"),
					resource.TestCheckResourceAttrPair(dataSourceName, "master_user_secret_arn", resourceName, "master_user_secret_arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "network_type", resourceName, "network_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags.Name", resourceName, "tags.Name"),
				),
			},
		},
	})
}

func testAccClusterDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
resource "aws_db_subnet_group" "test" {
  name       = %[1]q
  subnet_ids = aws_subnet.test[*].id
}

resource "aws_rds_cluster" "test" {
  cluster_identifier              = %[1]q
  database_name                   = "test"
  db_cluster_parameter_group_name = "default.aurora5.6"
  db_subnet_group_name            = aws_db_subnet_group.test.name
  master_password                 = "avoid-plaintext-passwords"
  master_username                 = "tfacctest"
  skip_final_snapshot             = true

  tags = {
    Name = %[1]q
  }
}

data "aws_rds_cluster" "test" {
  cluster_identifier = aws_rds_cluster.test.cluster_identifier
}
`, rName))
}

func testAccClusterDataSourceConfig_manage_password(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
resource "aws_db_subnet_group" "test" {
  name       = %[1]q
  subnet_ids = aws_subnet.test[*].id
}

resource "aws_rds_cluster" "test" {
  cluster_identifier              = %[1]q
  database_name                   = "test"
  db_cluster_parameter_group_name = "default.aurora5.6"
  db_subnet_group_name            = aws_db_subnet_group.test.name
  master_username                 = "tfacctest"
  manage_master_user_password     = true
  skip_final_snapshot             = true

  tags = {
    Name = %[1]q
  }
}

data "aws_rds_cluster" "test" {
  cluster_identifier = aws_rds_cluster.test.cluster_identifier
}
`, rName))
}
