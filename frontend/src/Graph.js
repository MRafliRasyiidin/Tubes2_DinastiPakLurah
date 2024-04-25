import React, { useState, useEffect } from 'react';
import Graph from 'react-vis-network-graph';
import { v4 as uuidv4 } from 'uuid';

function NodeGraph({ darkmode, start, target }) {
  const [graph, setGraph] = useState({ nodes: [], edges: [] });

  useEffect(() => {
    const handleSearch = () => {
      const nodes = fetchNodeList(start, target);
      generateGraph(nodes);
    };

    const fetchNodeList = (start, target) => {
      const nodeList = ['Node 1', 'Node 2', 'Node 3', 'Node 4', 'Node 5'];

      nodeList.push(target);
      nodeList.unshift(start);

      return nodeList;
    };

    const generateGraph = (nodes) => {
      const graphData = {
        nodes: [],
        edges: []
      };

      for (let i = 0; i < nodes.length; i++) {
        graphData.nodes.push({
          id: `${uuidv4()}`, // Change node ID to string
          label: `link${i}`,
          title: nodes[i]
        });

        if (i < nodes.length - 1) {
          graphData.edges.push({
            from: `${i}`, // Change from ID to string
            to: `${i + 1}`, // Change to ID to string
            id: `${uuidv4()}` // Change edge ID to string
          });
        }
      }
      setGraph(graphData);
    };

    if (start && target) {
      handleSearch();
    }
  }, [start, target]);

  const events = {
    select: function (event) {
      var { nodes, edges } = event;
      console.log(edges);
      console.log(nodes);
    }
  };

  console.log(graph)
  return (
    <div>
      <div style={{ width: '100vh', height: '75vh' }} className="flex items-center align-middle justify-center border border-red-500 mt-40 mb-10">
        <Graph
          graph={graph}
          events={events}
          options={{ layout: { hierarchical: false }, physics: false }}
        />
      </div>
    </div>
  );
}

export default NodeGraph;
