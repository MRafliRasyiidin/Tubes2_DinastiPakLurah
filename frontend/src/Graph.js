import React, { useState, useEffect } from 'react';
import Graph from 'react-vis-network-graph';
import { v4 as uuidv4 } from 'uuid';

function NodeGraph({ darkmode, start, target, listSolution}) {
  const [graph, setGraph] = useState({ nodes: [], edges: [] });
  const [length, setLength] = useState(0);

  useEffect(() => {
    const handleSearch = async () => {
      const nodes = await fetchNodeList(start, target);
      generateGraph(nodes);
    };

    const fetchNodeList = async (start, target, listSolution) => {
      // template (ini dihilangin nanti)
      const nodeLists = [
        ['Node 1', 'Node 2', 'Node 3'],
        ['node1', 'node2', 'node3'], 
        ['waduh', 'kuda', 'anjing'],
      ];
      // ini dihilangin kalo udah ada yang fix
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
          setLength(path.length);
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
      <div className = "flex flex-col items-center gap-1 mt-40">
        <div className = "flex items-center justify-center border ">
          Found in Depth: {length-1} 
        </div>
        <div style={{ width: '100vh', height: '75vh' }} className=" items-center align-middle justify-center border  bg-slate-300 rounded-xl mb-10">
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
