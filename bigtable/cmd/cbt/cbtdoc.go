// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// DO NOT EDIT. THIS IS AUTOMATICALLY GENERATED.
// Run "go generate" to regenerate.
//go:generate go run cbt.go -o cbtdoc.go doc

/*
Cbt is a tool for doing basic interactions with Cloud Bigtable. To learn how to
install the cbt tool, see the
[cbt overview](https://cloud.google.com/bigtable/docs/go/cbt-overview).

Usage:

	cbt [options] command [arguments]

The commands are:

	count                     Count rows in a table
	createinstance            Create an instance with an initial cluster
	createcluster             Create a cluster in the configured instance (replication alpha)
	createfamily              Create a column family
	createtable               Create a table
	updatecluster             Update a cluster in the configured instance
	deleteinstance            Deletes an instance
	deletecluster             Deletes a cluster from the configured instance (replication alpha)
	deletecolumn              Delete all cells in a column
	deletefamily              Delete a column family
	deleterow                 Delete a row
	deletetable               Delete a table
	doc                       Print godoc-suitable documentation for cbt
	help                      Print help text
	listinstances             List instances in a project
	listclusters              List instances in an instance
	lookup                    Read from a single row
	ls                        List tables and column families
	mddoc                     Print documentation for cbt in Markdown format
	read                      Read rows
	set                       Set value of a cell
	setgcpolicy               Set the GC policy for a column family
	waitforreplication        Blocks until all the completed writes have been replicated to all the clusters (replication alpha)
	version                   Print the current cbt version

Use "cbt help <command>" for more information about a command.

The options are:

	-project string
		project ID, if unset uses gcloud configured project
	-instance string
		Cloud Bigtable instance
	-creds string
		if set, use application credentials in this file


Alpha features are not currently available to most Cloud Bigtable customers. The
features might be changed in backward-incompatible ways and are not recommended
for production use. They are not subject to any SLA or deprecation policy.

For convenience, values of the -project, -instance, -creds,
-admin-endpoint and -data-endpoint flags may be specified in
/usr/local/google/home/igorbernstein/.cbtrc in this format:
	project = my-project-123
	instance = my-instance
	creds = path-to-account-key.json
	admin-endpoint = hostname:port
	data-endpoint = hostname:port
All values are optional, and all will be overridden by flags.



Count rows in a table

Usage:
	cbt count <table>




Create an instance with an initial cluster

Usage:
	cbt createinstance <instance-id> <display-name> <cluster-id> <zone> <num-nodes> <storage type>
	  instance-id					Permanent, unique id for the instance
	  display-name	  			Description of the instance
	  cluster-id						Permanent, unique id for the cluster in the instance
	  zone				  				The zone in which to create the cluster
	  num-nodes	  				The number of nodes to create
	  storage-type					SSD or HDD





Create a cluster in the configured instance (replication alpha)

Usage:
	cbt createcluster <cluster-id> <zone> <num-nodes> <storage type>
	  cluster-id		Permanent, unique id for the cluster in the instance
	  zone				  The zone in which to create the cluster
	  num-nodes	  The number of nodes to create
	  storage-type	SSD or HDD





Create a column family

Usage:
	cbt createfamily <table> <family>




Create a table

Usage:
	cbt createtable <table> [families=family[:(maxage=<d> | maxversions=<n>)],...] [splits=split,...]
	  families: Column families and their associated GC policies. See "setgcpolicy".
	  					 Example: families=family1:maxage=1w,family2:maxversions=1
	  splits:   Row key to be used to initially split the table




Update a cluster in the configured instance

Usage:
	cbt updatecluster <cluster-id> [num-nodes=num-nodes]
	  cluster-id		Permanent, unique id for the cluster in the instance
	  num-nodes		The number of nodes to update to




Deletes an instance

Usage:
	cbt deleteinstance <instance>




Deletes a cluster from the configured instance (replication alpha)

Usage:
	cbt deletecluster <cluster>




Delete all cells in a column

Usage:
	cbt deletecolumn <table> <row> <family> <column> [app-profile=<app profile id>]
	  app-profile=<app profile id>		The app profile id to use for the request (replication alpha)





Delete a column family

Usage:
	cbt deletefamily <table> <family>




Delete a row

Usage:
	cbt deleterow <table> <row> [app-profile=<app profile id>]
	  app-profile=<app profile id>		The app profile id to use for the request (replication alpha)





Delete a table

Usage:
	cbt deletetable <table>




Print godoc-suitable documentation for cbt

Usage:
	cbt doc




Print help text

Usage:
	cbt help [command]




List instances in a project

Usage:
	cbt listinstances




List instances in an instance

Usage:
	cbt listclusters




Read from a single row

Usage:
	cbt lookup <table> <row> [app-profile=<app profile id>]
	  app-profile=<app profile id>		The app profile id to use for the request (replication alpha)





List tables and column families

Usage:
	cbt ls			List tables
	cbt ls <table>		List column families in <table>




Print documentation for cbt in Markdown format

Usage:
	cbt mddoc




Read rows

Usage:
	cbt read <table> [start=<row>] [end=<row>] [prefix=<prefix>] [regex=<regex>] [count=<n>] [app-profile=<app profile id>]
	  start=<row>		Start reading at this row
	  end=<row>		Stop reading before this row
	  prefix=<prefix>	Read rows with this prefix
	  regex=<regex> 	Read rows with keys matching this regex
	  count=<n>		Read only this many rows
	  app-profile=<app profile id>		The app profile id to use for the request (replication alpha)





Set value of a cell

Usage:
	cbt set <table> <row> [app-profile=<app profile id>] family:column=val[@ts] ...
	  app-profile=<app profile id>		The app profile id to use for the request (replication alpha)
	  family:column=val[@ts] may be repeated to set multiple cells.

	  ts is an optional integer timestamp.
	  If it cannot be parsed, the `@ts` part will be
	  interpreted as part of the value.




Set the GC policy for a column family

Usage:
	cbt setgcpolicy <table> <family> ( maxage=<d> | maxversions=<n> )

	  maxage=<d>		Maximum timestamp age to preserve (e.g. "1h", "4d")
	  maxversions=<n>	Maximum number of versions to preserve




Blocks until all the completed writes have been replicated to all the clusters (replication alpha)

Usage:
	cbt waitforreplication <table>




Print the current cbt version

Usage:
	cbt version




*/
package main
