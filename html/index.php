<?php  
	
	$graphContent = file_get_contents( "graph.json" );
	$decodedJson = (json_decode($graphContent, true));


?>


<!DOCTYPE html>
<html>
<head>



	<title>Project 4</title>
	
	<link href="vis.css" rel="stylesheet" type="text/css" />
	<script type="text/javascript" src="vis.js"></script>
</head>
<body>

<h1>Goject 4 - a GO-netsim     CurrentTimeStamp: <?php echo $decodedJson['Name'] ;?></h1>



<div>
	<br><br>
<?php 

	$idToIP = array();

  	$id = 1 ;
	foreach ($decodedJson['Nodes'] as $key => $value) {
		$idToIP[$key] = $id;
		$id = $id + 1;
	}





 ?>

</div>

<div id="mynetwork" style="height: 600px;width: 100%;"></div>

</body>
	
	<script type="text/javascript">
		
  // create an array with nodes
  var nodes = [

  	<?php



  	$id = 1 ;
	foreach ($decodedJson['Nodes'] as $key => $value) {
		echo "{id: ".$id.", label: '".$key."'},";
		$id = $id + 1;
	}


?>


  ];

  // create an array with edges
  var edges = [
<?php

	foreach ($decodedJson['Nodes'] as $key => $value) {
		foreach ($value['Edges'] as $keyy => $value) {
			echo "{from: ".$idToIP[$key].", to: ".$idToIP[$keyy].", label: '".$value."',     font: {align: 'bottom'}, length:200},";
		}
	
	}

?>


  ];

  // create a network
  var container = document.getElementById('mynetwork');
  var data = {
    nodes: nodes,
    edges: edges
  };
  var options = {};
  var network = new vis.Network(container, data, options);

	</script>



</html>