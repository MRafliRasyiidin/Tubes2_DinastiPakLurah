import React, { useState, useEffect } from 'react';
import Graph from 'react-graph-vis';

function NodeGraph({ darkmode, start, target }) {
  const [graph, setGraph] = useState({ nodes: [], edges: [] });

//   console.log(start); 
//   console.log(target);
  const handleSearch = () => {
    // Logic to fetch node list from start to target
    const nodes = fetchNodeList(start, target);
    generateGraph(nodes);
  };

  useEffect(() => {
    if (start && target) {
      handleSearch();
    }
  }, [start, target]);

  const fetchNodeList = (start, target) => {
    // Example node list, replace this with your actual data fetching logic
    const nodeList = ['Node 1', 'Node 2', 'Node 3', 'Node 4', 'Node 5'];
    const startIndex = nodeList.indexOf(start);
    const targetIndex = nodeList.indexOf(target);

    if (startIndex === -1 || targetIndex === -1 || startIndex >= targetIndex) {
      return [];
    }

    return nodeList.slice(startIndex, targetIndex + 1);
  };

  const generateGraph = (nodes) => {
    const graphData = {
      nodes: nodes.map((node, index) => ({ id: index, label: node })),
      edges: nodes.slice(0, -1).map((_, index) => ({ from: index, to: index + 1 })),
    };
    setGraph(graphData);
  };

  return (
    <div>
      {start && target && (
        <div style={{ width: '80%', height: '400px' }}>
          <Graph
            graph={graph}
            options={{ layout: { hierarchical: false }, physics: false }}
          />
        </div>
      )}
    </div>
  );
}

export default NodeGraph;
