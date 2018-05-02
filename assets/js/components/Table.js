'use strict';

import { h, render, Component } from 'preact';
import * as numbers from '../lib/numbers.js';
import Client from '../lib/client.js';
import { bind } from 'decko';

const dayInSeconds = 60 * 60 * 24;

class Table extends Component {

  constructor(props) {
    super(props)

    this.state = {
      records: [],
      limit: 15,
      loading: true,
      before: props.before,
      after: props.after,
    }
  }

  componentDidMount() {
      this.fetchRecords()
  }

  componentWillReceiveProps(newProps, prevState) {
      if(newProps.before == prevState.before && newProps.after == prevState.after) {
        return;
      }

      this.setState({
        before: newProps.before,
        after: newProps.after,
      });
      this.fetchRecords();
  }

  @bind
  fetchRecords() {
    this.setState({ loading: true });
  
    Client.request(`${this.props.endpoint}?before=${this.state.before}&after=${this.state.after}&limit=${this.state.limit}`)
      .then((d) => {
        this.setState({ 
          loading: false,
          records: d,
        });
      });
  }

  render(props, state) {
    const tableRows = state.records !== null ? state.records.map((p, i) => {
      let ahref = document.createElement('a'); ahref.href = p.value;
      let classes = "table-row w" + Math.round(p.percentage_of_total);
      let label = ahref.pathname;
      if( props.showHostname ) {
        label = ahref.hostname.replace('www.', '') + (ahref.pathname.length > 1 ? ahref.pathname : '');
      }

      return(
      <div class={classes}>
        <div class="cell main-col"><a href={ahref.href}>{label}</a></div>
        <div class="cell">{p.count}</div>
        <div class="cell">{p.count_unique||"-"}</div>           
      </div>
    )}) : <div class="table-row">Nothing here, yet.</div>;

    const loadingOverlay = state.loading ? <div class="loading-overlay"><div></div></div> : '';

    return (
      <div class="box box-pages animated fadeInUp delayed_04s">
            
        <div class="table-row header">
          {props.headers.map((header, i) => {
            let classes = i === 0 ? 'main-col cell' : 'cell';
            return (<div class={classes}>{header}</div>) 
            })}        
        </div>

       {tableRows}
      </div>
    )
  }
}

export default Table
