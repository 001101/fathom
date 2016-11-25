'use strict';

import { h, render, Component } from 'preact';
import Chart from 'chart.js'

Chart.defaults.global.tooltips.xPadding = 10;
Chart.defaults.global.tooltips.yPadding = 10;
Chart.defaults.global.layout = { padding: 10 }

class Graph extends Component {
  constructor(props) {
    super(props)

    this.state = {
      visitorData: [],
      pageviewData: []
    }
    this.fetchData = this.fetchData.bind(this);
    this.fetchData(props.period);
  }

  componentWillReceiveProps(newProps) {
    if(this.props.period != newProps.period) {
      this.fetchData(newProps.period)
    }
  }

  refreshChart() {
    if( ! this.canvas ) { return; }

    // clear canvas
    var newCanvas = document.createElement('canvas');
    this.canvas.parentNode.style.minHeight = this.canvas.parentNode.clientHeight + "px";
    this.canvas.parentNode.replaceChild(newCanvas, this.canvas);
    this.canvas = newCanvas;

    if( this.chart ) { this.chart.clear(); }
    this.chart = new Chart(this.canvas, {
      type: 'line',
      data: {
        labels: this.state.visitorData.map((d) => d.Label),
        datasets: [
          {
            label: '# of Visitors',
            data: this.state.visitorData.map((d) => d.Count),
            backgroundColor: 'rgba(255, 155, 0, .6)',
            pointStyle: 'rect',
            pointBorderWidth: 0.1,
          },
          {
            label: '# of Pageviews',
            data: this.state.pageviewData.map((d) => d.Count),
            backgroundColor: 'rgba(0, 155, 255, .4)',
            pointStyle: 'rect',
            pointBorderWidth: 0.1,
          }
      ],
    }
    });
  }

  fetchData(period) {
    // fetch visitor data
    fetch('/api/visits/count/day?period=' + period, {
      credentials: 'include'
    }).then((r) => {
      if( r.ok ) {
        return r.json();
      }
      throw new Error();
    }).then((data) => {
      this.setState({ visitorData: data })
      window.requestAnimationFrame(this.refreshChart.bind(this));
    });

    // fetch pageview data
    fetch('/api/pageviews/count/day?period=' + period, {
      credentials: 'include'
    }).then((r) => {
      if( r.ok ) {
        return r.json();
      }

      throw new Error();
    }).then((data) => {
        this.setState({ pageviewData: data })
        window.requestAnimationFrame(this.refreshChart.bind(this));
    });
  }

  render() {
    return (
      <div class="block">
        <canvas height="100" ref={(el) => { this.canvas = el; }} />
      </div>
    )
  }
}

export default Graph
