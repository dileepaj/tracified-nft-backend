package configs

func GetMapHead() string {
	var maphead = `<html>
	<head>
	<meta charset="utf-8" />
	  <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.3/dist/leaflet.css"
		integrity="sha256-kLaT2GOSpHechhsozzB+flnD+zUyjE2LlfWPgU04xyI=" crossorigin="" />
	  <script src="https://unpkg.com/leaflet@1.9.3/dist/leaflet.js"
		integrity="sha256-WBkoXOwTeyKclOHuWtc+i2uENFpDZ9YPdf5Hf+D7ewM=" crossorigin=""></script>
	  <script src="https://unpkg.com/leaflet-polylinedecorator@1.6.0/dist/leaflet.polylineDecorator.js"
		crossorigin=""></script>
	  <script src='https://unpkg.com/leaflet-arc/bin/leaflet-arc.min.js'></script>
	
	  <style>
		#map {
			position: absolute;
			top: 0;
			bottom: 0;
			right: 0;
			left: 0;
		}
	  </style>
	</head>
	`
	return maphead
}

func GetMapBody() string {
	var mapBody = `
		<body>
			<div id="map"></div>
  		</body>
  	`
	return mapBody
}
func GetScriptStart() string {
	var scriptStart = `
	<script>
		var locations = [
	`
	return scriptStart
}

func GetLowerScript() string {
	var lowerScript = `];
	function getPointers() {
		var pointArr = [];
		for (var k = 0; k < locations.length; k++) {
		  var point = [locations[k][1], locations[k][2]];
		  pointArr.push(point);
		}
		return pointArr;
	  }
	
	  //create markers with description
	  var map = L.map("map");
	  mapLink = '<a href="https://openstreetmap.org">OpenStreetMap</a>';
	  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
		attribution: "&copy; " + mapLink + " Contributors",
		maxZoom: 18,
	  }).addTo(map);
	
	  for (var i = 0; i < locations.length; i++) {
		var marker = new L.marker([locations[i][1], locations[i][2]]).addTo(map);
		var popup = new L.popup({ autoClose: false })
		  .setLatLng([locations[i][1], locations[i][2]])
		  .setContent(` + "`<p><center><b>${locations[i][0]}</b></center></p>`" + `)
		  .openOn(map);
	
		marker.bindPopup(popup);
	  }
	
	  //add polylines with arrowheads
	  for (var j = 0; j < locations.length - 1; j++) {
		var polyline = L.Polyline.Arc([locations[j][1], locations[j][2]], [locations[j + 1][1], locations[j + 1][2]]);
		L.polylineDecorator(polyline, {
		  patterns: [
			{
			  offset: 10,
			  repeat: 20,
			  symbol: L.Symbol.dash({
				pixelSize: 10,
				pathOptions: {
				  color: "#000",
				  weight: 2,
				},
			  }),
			},
			{
			  offset: "100%",
			  repeat: 0,
			  symbol: L.Symbol.arrowHead({
				pixelSize: 20,
				polygon: false,
				pathOptions: {
				  stroke: true,
				  color: "#000",
				},
			  }),
			},
		  ],
		}).addTo(map);
	  }
	
	  map.fitBounds(getPointers(), { maxZoom: 4 });
	</script>
	
	</html>
	`
	return lowerScript
}
