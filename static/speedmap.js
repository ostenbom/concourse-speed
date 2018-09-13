function speedMap() {
  var map = {};

  map.render = function() {
    var svg = d3.select('#speedmap').append('svg')
                .attr('width', '100%')
                .attr('height', '100%')
                .append('g')
                .attr('transform', 'translate(0,0)');
  };

  return map;
}
