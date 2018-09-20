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

describe("the data", function() {
  var exampleData;

  beforeEach(function() {
    exampleData = JSON.parse(exampleDataString);
  });

  it("converts server data to map data", function() {
    parsedEvents = parseBuildEventTimes(exampleData);
    mapData = buildEventsToMapData(exampleData, 10, maxInBox);
    console.log(mapData)
    expect(mapData[0].Job).toBe(0);
    expect(mapData[1].Job).toBe(1);
    expect(mapData[mapData.length - 1].Job).toBe(4);
  });
});


var exampleDataString = `
[
{
"Build":"26",
"Job":"guardian-containerd-xenial",
"Status":"succeeded",
"Start":"2018-09-13T16:25:23.107323+01:00",
"End":"2018-09-13T16:39:37.263175+01:00"
},
{
"Build":"3464",
"Job":"guardian-windows-periodic",
"Status":"succeeded",
"Start":"2018-09-13T09:58:44.415586+01:00",
"End":"2018-09-13T10:00:40.267257+01:00"
},
{
"Build":"3468",
"Job":"guardian-xenial-periodic",
"Status":"succeeded",
"Start":"2018-09-13T11:19:05.357905+01:00",
"End":"2018-09-13T11:44:26.442093+01:00"
},
{
"Build":"3466",
"Job":"guardian-trusty-periodic",
"Status":"succeeded",
"Start":"2018-09-13T09:18:33.042908+01:00",
"End":"2018-09-13T09:44:24.295419+01:00"
},
{
"Build":"3467",
"Job":"guardian-trusty-periodic",
"Status":"succeeded",
"Start":"2018-09-13T09:58:44.293686+01:00",
"End":"2018-09-13T10:24:30.848283+01:00"
},
{
"Build":"3469",
"Job":"guardian-trusty-periodic",
"Status":"succeeded",
"Start":"2018-09-13T11:19:05.283632+01:00",
"End":"2018-09-13T11:44:39.113786+01:00"
},
{
"Build":"3465",
"Job":"guardian-xenial-periodic",
"Status":"succeeded",
"Start":"2018-09-13T09:18:33.113619+01:00",
"End":"2018-09-13T09:43:54.244709+01:00"
},
{
"Build":"3466",
"Job":"grootfs-xenial-periodic",
"Status":"succeeded",
"Start":"2018-09-13T09:58:44.529213+01:00",
"End":"2018-09-13T10:01:01.267534+01:00"
},
{
"Build":"3471",
"Job":"guardian-trusty-periodic",
"Status":"succeeded",
"Start":"2018-09-13T12:39:13.330466+01:00",
"End":"2018-09-13T13:05:01.233263+01:00"
},
{
"Build":"3463",
"Job":"guardian-windows-periodic",
"Status":"succeeded",
"Start":"2018-09-13T09:18:33.185654+01:00",
"End":"2018-09-13T09:20:29.910194+01:00"
},
{
"Build":"3470",
"Job":"guardian-xenial-periodic",
"Status":"succeeded",
"Start":"2018-09-13T12:39:13.403216+01:00",
"End":"2018-09-13T13:04:12.959186+01:00"
},
{
"Build":"3475",
"Job":"guardian-trusty-periodic",
"Status":"succeeded",
"Start":"2018-09-13T15:19:47.513394+01:00",
"End":"2018-09-13T15:45:34.31831+01:00"
}
]
`;
