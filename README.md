# Go Graph Seeder
<img src="https://github.com/user-attachments/assets/b240a681-97c9-4979-b222-e3eaa8807a57" alt="go graph seeder" width="500" height="500">

### Requirements
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

### Run it!
Once the template.json is all set up **run**
```bash
go run main.go
```
