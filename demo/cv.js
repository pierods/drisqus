/*
<script src="https://d3js.org/d3-color.v1.min.js"></script>
<script src="https://d3js.org/d3-interpolate.v1.min.js"></script>
<script src="https://d3js.org/d3-scale-chromatic.v1.min.js"></script>

*/

function makeCV(dataURL) {
  d3.json(dataURL, drawCV)
}

function drawCV(data) {
  // get unique member names
  var membersO = {}
  data.forEach(function(play) {
    membersO[play[0]] = "-"
    membersO[play[1]] = "-"
  })
  //make member array
  var members = []
  Object.keys(membersO).forEach(function(member, index) {
    members.push(member)
  })

  // assign member names to colors
  //var colorPalette = d3.scaleOrdinal(d3.schemeCategory20c);
  scale = d3.scaleLinear().domain([0, members.length])
  var colors = {}
  members.forEach(function(member, index) {
    //colors[member] = colorPalette(index)
    colors[member] = d3.interpolateSpectral(scale(index))
  })

  // sort lexicographically
  var sortOrder = members.sort(d3.ascending)

  function sort(a, b) {
    return d3.ascending(sortOrder.indexOf(a), sortOrder.indexOf(b));
  }

  var ch = viz.ch().data(data)
    .padding(.01)
    .sort(sort)
    .innerRadius(430)
    .outerRadius(450)
    .duration(1000)
    .chordOpacity(0.3)
    .labelPadding(.03)
    .fill(function(d) {
      return colors[d];
    });

  var width = 1200,
    height = 1100;
  var svg = d3.select("body").append("svg").attr("height", height).attr("width", width);

  svg.append("g").attr("transform", "translate(600,550)").call(ch);

  // adjust height of frame in bl.ocks.org
  //d3.select(self.frameElement).style("height", height+"px").style("width", width+"px");
}
