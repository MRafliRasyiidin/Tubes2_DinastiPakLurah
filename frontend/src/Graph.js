import React, { useState, useEffect } from 'react';
import Graph from 'react-vis-network-graph';
import { v4 as uuidv4 } from 'uuid';

function NodeGraph({ darkmode, start, target }) {
  const [graph, setGraph] = useState({ nodes: [], edges: [] });

  useEffect(() => {
    const handleSearch = async () => {
      const nodes = await fetchNodeList(start, target);
      generateGraph(nodes);
    };

    const fetchNodeList = async (start, target) => {
      // template
      const nodeList = ['Node 1', 'Node 2', 'Node 3', 'Node 4', 'Node 5'];

      nodeList.push(target);
      nodeList.unshift(start);

      return nodeList;
    };

    const generateGraph = (nodes) => {
      var graphData = {
        nodes: [],
        edges: []
      };

      for (let i = 0; i < nodes.length; i++) {
        graphData.nodes.push({
          id: `${uuidv4()}`,
          label: nodes[i],
          title: nodes[i],
        });

        if(i >= 1){
            if (i < nodes.length) {
              graphData.edges.push({
                from: graphData.nodes[i-1].id,
                to: graphData.nodes[i].id,
                id: `${uuidv4()}`,
              });
            }
        }
      }
      setGraph(graphData);
    };

    if (start && target) {
      handleSearch();
    }
  }, [start, target]);

  const options = {
    edges: {
      color: {
        color: '#b32e2e',
        highlight: '#b32e2e',
        hover: '#b32e2e',
        opacity: 1,
        inherit: false,
      },
      width: 2,
      dashes: true,
    },
    height: '700px',
    layout: {
      hierarchical: false
    },
  };

  const events = {
    select: function (event) {
      var { nodes, edges } = event;
      console.log(nodes);
      console.log(edges);
    }
  };

  return (
    <div>
      <div style={{ width: '100vh', height: '75vh' }} className="flex items-center align-middle justify-center border border-red-500 mt-40 mb-10">
        <Graph
          key={uuidv4()}
          options={options}
          graph={graph}
          events={events}
        />
      </div>
    </div>
  );
}

export default NodeGraph;
