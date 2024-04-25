import React, { useState, useEffect } from 'react';
import Graph from 'react-vis-network-graph';
import { v4 as uuidv4 } from 'uuid';

function NodeGraph({ darkmode, start, target}) {
  const [graph, setGraph] = useState({ nodes: [], edges: [] });

  useEffect(() => {
    const handleSearch = async () => {
      const nodes = await fetchNodeList(start, target);
      generateGraph(nodes);
    };

    const fetchNodeList = async (start, target, listSolution) => {
      // template
      const nodeLists = [
        ['Node 1', 'Node 2', 'Node 3', 'Node 4', 'Node 5'],
        ['node1', 'node2', 'node3'], 
        ['waduh', 'kuda', 'anjing'],
      ];
      for(let node of nodeLists){
        node.push(target);
        node.unshift(start);
      }

      console.log(nodeLists);
      return nodeLists;
    };

    const generateGraph = (paths) => {
      var graphData = {
        nodes: [],
        edges: []
      };

      for(let path of paths){
        for (let i = 0; i < path.length; i++) {
          const node = path[i]
          const isNewNode = !graphData.nodes.find(n => n.label === node);

          if(isNewNode){
            graphData.nodes.push({
              id: `${uuidv4()}`,
              label: path[i],
              title: path[i],
            });
          }
  
          if(i > 0){
                const fromNode = path[i - 1];
                const toNode = path[i];
                const fromNodeObj = graphData.nodes.find(n => n.label === fromNode);
                const toNodeObj = graphData.nodes.find(n => n.label === toNode);
                if(fromNodeObj && toNodeObj){
                  graphData.edges.push({
                    from: fromNodeObj.id,
                    to: toNodeObj.id,
                    id: `${uuidv4()}`,
                  });
              }
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
      dashes: false,
    },
    height: '700px',
    layout: {
      hierarchical: false,
      improvedLayout: true,
    },
    physics: true, 

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
      <div style={{ width: '100vh', height: '75vh' }} className="flex items-center align-middle justify-center border  bg-slate-300 rounded-xl  mt-40 mb-10">
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
