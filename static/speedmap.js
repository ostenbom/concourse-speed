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

function maxInBox(boxData) {
  var maxData = {};
  maxData.Time = 0;

  for (var i = 0; i < boxData.length; i++) {
    var data = boxData[i];
    if (data.Time > maxData.Time) {
      maxData = data;
    }
  }

  return maxData.Time
}

function getLatestStart(events) {
  var latest = new Date('December 17, 1995 03:24:00');
  for (var i = 0; i < events.length; i++) {
    var event = events[i];
    if (event.Start > latest) {
      latest = event.Start;
    }
  }

  return latest
}

function getEarliestStart(events) {
  var earliest = Date.now();
  for (var i = 0; i < events.length; i++) {
    var event = events[i];
    if (event.Start < earliest) {
      earliest = event.Start;
    }
  }

  return earliest
}

function getJobNames(events) {
  var jobNames = [];
  for (var i = 0; i < events.length; i++) {
    var event = events[i];
    if (!jobNames.includes(event.Job)) {
      jobNames.push(event.Job);
    }
  }

  return jobNames
}

function getEventsInBox(events, boxTime, boxTimeLength, jobName) {
  var boxData = [];
  for (var eventIndex = 0; eventIndex < events.length; eventIndex++) {
    var event = events[eventIndex];
    // console.log(event.Start.getTime(), boxTime, boxTime + boxTimeLength, jobName);
    if (
      event.Start.getTime() >= boxTime &&
      event.Start.getTime() < (boxTime + boxTimeLength) &&
      event.Job == jobName
    ) {
      // Warning optimisation do deepcopy
      events.splice(eventIndex, 1);

      var matchingEvent = {};

      matchingEvent.Build = event.Build;
      matchingEvent.Job = event.Job;
      matchingEvent.Status = event.Status;
      matchingEvent.Time = (event.End - event.Start) / 1000;

      boxData.push(matchingEvent);
    }

  }

  return boxData
}

function buildEventsToMapData(trueEvents, numBoxes, boxChooser) {
  var events = trueEvents.slice(0);
  var maxTime = getLatestStart(events);
  var minTime = getEarliestStart(events);

  var totalTime = (maxTime - minTime) / 1000;
  var boxTimeLength = (totalTime / numBoxes) * 1000;

  var jobNames = getJobNames(events);
  var numJobs = jobNames.length;

  var mapData = [];

  var job, box;
  for (job = 0; job < numJobs; job++) {
    var boxTime = minTime.getTime();
    for (box = 0; box < numBoxes; box++) {
      var jobName = jobNames[job];
      var eventsInBox = getEventsInBox(events, boxTime, boxTimeLength, jobName);
      var boxValue = boxChooser(eventsInBox);
      var boxData = {};

      boxData.Job = job;
      boxData.Box = box;
      boxData.Value = boxValue;

      mapData.push(boxData);

      boxTime += boxTimeLength;
    }
  }

  return mapData
}

function parseBuildEventTimes(events) {
  for (var i = 0; i < events.length; i++) {
    var event = events[i]
    event.Start = new Date(event.Start);
    event.End = new Date(event.End);
  }

  return events
}

function getJson(url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'json';
    xhr.onload = function() {
      var status = xhr.status;
      if (status === 200) {
        callback(null, xhr.response);
      } else {
        callback(status, xhr.response);
      }
    };
    xhr.send();
}
