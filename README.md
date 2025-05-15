# Go Graph Seeder
<img src="https://github.com/user-attachments/assets/b240a681-97c9-4979-b222-e3eaa8807a57" alt="go graph seeder" width="500" height="500">

Welcome to Go Graph Seeder! This handy application is designed to help you easily create and manage nodes and edges in a Gremlin-based graph database. Whether you’re a developer, data scientist, or just someone curious about graph databases, this tool is here to make your life easier!
### Why Choose Go Graph Seeder?
- **Connect with Ease**: The application connects seamlessly to a Gremlin server running at localhost:8182. Just fire it up, and you’re ready to go!
- **Customize Your Nodes**: With our JSON configuration file, you have the power to define both startNodes and endNodes. You can specify their properties, how many of each you want, and even how they relate to one another.
- **Diverse Property Types**: Need random data? No problem! Go Graph Seeder supports various property types like emails, names, addresses, and more. You can create realistic data in no time.
- **Fun with Enums**: Want to add some randomness? Use the enum property type to generate values from a set of options with weighted probabilities. It’s a great way to add variety!

![go_seeder_demo](https://github.com/user-attachments/assets/f01a6372-bc7e-46ba-80d9-95d1e1273510)

### Requirements to run the source code
- Go version >= 1.22.3

### Initial configuration
Rename the file ***template_example.json*** to ***template.json***

### template.json
JSON configuration file used for generating nodes and edges in a Gremlin-based graph database. The configuration defines two main sections: **startNodes** and **endNodes**. Each section outlines the properties of nodes, their counts, and any relationships or connections they have.

### JSON Configuration Structure
```json
{
  "startNodes": [
    {
      "nodeCount": 1,
      "relationships": [{ "edge": "has_state", "endNode": "state" }],
      "nodeProperties": {
        "label": "country",
        "country_name": "country",
        "created_at": "date",
        "updated_at": "date"
      }
    }
  ],
  "endNodes": [
    {
      "nodeCount": 1,
      "nodeProperties": {
        "label": "state",
        "state_name": "state",
        "created_at": "date"
      }
    }
  ]
}
```
Key Sections
1. **startNodes**
The **startNodes** array defines the primary nodes that will be created in the graph database. These nodes may have relationships to other nodes.
- **nodeCount**: Specifies the number of nodes to create. For example, 1 indicates that one node will be generated.
- **relationships**:
   - Defines edges that connect this node to other nodes in the graph.
   - Each relationship specifies:
     - **edge**: The label of the edge (e.g., "**has_state**").
     - **endNode**: The label of the target node to connect to (e.g., "**state**").

- **nodeProperties**:
  - Contains the properties for each node. Each property is defined as a key-value pair where the key is the property name, and the value specifies the type of data (e.g., "**country**", "**date**").
  - **label**: Specifies the type of the node (e.g., "**country**").

  - **example**
```json
{
  "label": "country",
  "country_name": "country",
  "created_at": "date",
  "updated_at": "date"
}
```
Example of a **startNode**:

A **country** node with properties **country_name**, **created_at**, and **updated_at** will be created and connected to a state node with an edge labeled "has_state".

2. **endNodes**
The **endNodes** array defines secondary nodes that act as targets for relationships or standalone entities in the graph.
- **nodeCount**: Specifies the number of nodes to create. For example,** 1** indicates that one node will be generated.
- nodeProperties:
  - Contains the properties for each node, defined similarly to **startNodes**.
- **label**: Specifies the type of the node (e.g., "**state**").
-
example
```json
{
  "label": "state",
  "state": "state",
  "created_at": "date"
}
```
Example of an **endNode**:
A **state** node with properties **state_name** and **created_at** will be created.

**Example Graph Representation**

**Nodes**:
1. **country** node:
    - Properties
      - label:country
      - country_name:country
      - created_at:date
      - updated_at:date
2. **state** node:
   - Properties
      - label:state
      - state_name:state
      - created_at:date

**Edges**
- Edge Label: "**has_state**"
- Source Node: "**country**"
- Target Node: "**state**"

# Supported Property Types

| **Property Type**        | **Description**                                                                                       | **Example Output**                      |
|---------------------------|-------------------------------------------------------------------------------------------------------|------------------------------------------|
| `email`                  | Generates a random email address.                                                                    | `john.doe@example.com`                  |
| `fname`                  | Generates a random first name.                                                                       | `John`                                  |
| `lname`                  | Generates a random last name.                                                                        | `Doe`                                   |
| `address`                | Generates a random street address.                                                                   | `123 Elm St`                            |
| `city`                   | Generates a random city name.                                                                        | `Springfield`                           |
| `state`                  | Generates a random U.S. state name.                                                                  | `California`                            |
| `country`                | Generates a random country name.                                                                     | `United States`                         |
| `postal`                 | Generates a random postal code.                                                                      | `90210`                                 |
| `phone`                  | Generates a random phone number.                                                                     | `(555) 123-4567`                        |
| `job`                    | Generates a random job title.                                                                        | `Software Engineer`                     |
| `language`               | Generates a random language name.                                                                   | `English`                               |
| `date`                   | Generates a random date in the format `YYYY-MM-DDTHH:MM:SSZ`.                                       | `2024-11-22T14:53:00Z`                  |
| `timestamp`              | Generates the current UNIX timestamp.                                                               | `1700658753`                            |
| `size`                   | Generates a random numerical size between 1 and 1000.                                               | `523`                                   |
| `enum`                   | Selects a value from a predefined set of weighted options.                                           | Based on input values (see below).      |

**enum** Property Type

The **enum** type allows generating values from a predefined set, where each value has a weight determining its likelihood of being selected.
Input example
```json
{
  "type": "enum",
  "values": [
    { "value": "option1", "weight": 0.2 },
    { "value": "option2", "weight": 0.5 },
    { "value": "option3", "weight": 1 }
  ]
}
```
Behavior:

Values with higher weights are more likely to be selected.
- Example outputs:
  - **option3** appears most frequently.
  - **option1** appears least frequently.

### Sample template.json configuration
If a property is created with a undefined **propType** it will default to a random word. See example below, the **music_type** property.

```json
{
  "nodeProperties": {
    "label": "person",
    "email": "email",
    "first_name": "fname",
    "last_name": "lname",
    "music_type: "loren ipsum",
    "ccategory": {
      "type": "enum",
      "values": [
        { "value": "category1", "weight": 0.3 },
        { "value": "category2", "weight": 0.7 }
      ]
    }
  }
}
```
Generated output
```json
{
  "label": "person",
  "email": "jane.doe@example.com",
  "first_name": "Jane",
  "last_name": "Doe",
  "music_type": "random word",
  "category": "category2"
}
```
### Run it!
Once the template.json is all set up **run**
```bash
go run main.go

#got the binary
#the file template.json has to be in the same directory as the binary.
./go-graph-seeder

```

## TODO

- [ ] Make the host dynamic
- [ ] Make the port dynamic
- [ ] Make the file name & location dynamic
- [ ] Handle multiple levels of edges, right now it handles 1 level (ex. a -> b )
