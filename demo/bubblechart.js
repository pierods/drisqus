function makeBubbleChart(url) {
  console.log(url)
  d3.csv(url, preparseCSVRow, bubblechart);
}
// convert value to int, skip rows where value is absent
function preparseCSVRow(row) {
  // a row has been read and converted to
  // a struct, witht id and value fields
  // convert row.value to a number
  row.value = +row.value;
  // skip rows with nonnumber or absent value
  if (row.value)
    return row;
}
// gets called after whole csv has been parsed.
// error is an error, dataArray is an array of {id, value} structs
function bubblechart(error, dataArray) {
  if (error) throw error;

  var svg = d3.select("svg"),
    width = +svg.attr("width"),
    height = +svg.attr("height");

  var format = d3.format(",d");

  var color = d3.scaleOrdinal(d3.schemeCategory20c);

  var pack = d3.pack()
    .size([width, height])
    .padding(2);

  // Constructs a root node from the specified hierarchical data.
  //The specified data must be an object representing the root node.
  // root node means whole tree
  var root = d3.hierarchy({
      children: dataArray
    })
    .sum(function(d) {
      return d.value;
    })
    .each(function(d) {
      if (id = d.data.id) {
        var id, i = id.lastIndexOf(".");
        d.id = id;
        d.package = id.slice(0, i);
        d.class = id.slice(i + 1);
      }
    });

  var node = svg.selectAll(".node")
    .data(pack(root).leaves())
    .enter().append("g")
    .attr("class", "node")
    .attr("transform", function(d) {
      return "translate(" + d.x + "," + d.y + ")";
    });

  node.append("circle")
    .attr("id", function(d) {
      return d.id;
    })
    .attr("r", function(d) {
      return d.r;
    })
    .style("fill", function(d) {
      return color(d.package);
    });

  node.append("clipPath")
    .attr("id", function(d) {
      return "clip-" + d.id;
    })
    .append("use")
    .attr("xlink:href", function(d) {
      return "#" + d.id;
    });

  node.append("text")
    .attr("clip-path", function(d) {
      return "url(#clip-" + d.id + ")";
    })
    .selectAll("tspan")
    .data(function(d) {
      return d.class.split(/(?=[A-Z][^A-Z])/g);
    })
    .enter().append("tspan")
    .attr("x", 0)
    .attr("y", function(d, i, nodes) {
      return 13 + (i - nodes.length / 2 - 0.5) * 10;
    })
    .text(function(d) {
      return d;
    });

  node.append("title")
    .text(function(d) {
      return d.id + "\n" + format(d.value);
    });
}
