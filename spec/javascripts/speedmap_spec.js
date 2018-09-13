describe("the speedmap", function() {
  var chart;
  var chartDiv;

  beforeEach(function() {
    chartDiv = document.createElement("div");
    chartDiv.id = "speedmap";
    document.body.appendChild(chartDiv);

    chart = speedMap();
    chart.render();
  });

  afterEach(function() {
    d3.selectAll("svg").remove();
    document.body.removeChild(chartDiv);
  });

  it("should be created", function() {
    expect(getSvg().empty()).not.toBe(true);
  });

  it("should have width 100%", function() {
    expect(getSvg().attr('width')).toBe('100%');
  });

  it("should have height 100%", function() {
    expect(getSvg().attr('height')).toBe('100%');
  });

  function getSvg() {
    return d3.select("svg");
  }
});
