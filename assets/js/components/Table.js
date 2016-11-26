'use strict';

import { h, render, Component } from 'preact';

class Table extends Component {

  constructor(props) {
    super(props)

    this.state = {
      records: []
    }

    this.tableHeaders = props.headers.map(heading => <th>{heading}</th>);
    this.fetchRecords = this.fetchRecords.bind(this);
    this.fetchRecords(props.period);
  }

  labelCell(p) {
    if( this.props.labelCell ) {
      return this.props.labelCell(p);
    }

    return (
      <td>{p.Label}</td>
    )
  }

  componentWillReceiveProps(newProps) {
    if(this.props.period != newProps.period) {
      this.fetchRecords(newProps.period)
    }
  }

  fetchRecords(period) {
    return fetch('/api/'+this.props.endpoint+'?period=' + period, {
      credentials: 'include'
    }).then((r) => {
      if( r.ok ) {
        return r.json();
      }

      // TODO: Make this pretty.
      if( r.status == 401 ) {
        this.props.onAuthError();
      }

      // TODO: do something with error
      throw new Error();
    }).then((data) => {
      this.setState({ records: data })
    }).catch((e) => {

    });
  }

  render() {
    const tableRows = this.state.records.map((p, i) => (
      <tr>
        <td class="muted">{i+1}</td>
        {this.labelCell(p)}
        <td>{p.Count}</td>
        <td>{Math.round(p.Percentage)}%</td>
      </tr>
    ));

    return (
      <div class="block">
        <h3>{this.props.title}</h3>
        <table>
          <thead>
            <tr>{this.tableHeaders}</tr>
          </thead>
          <tbody>
            {tableRows}
          </tbody>
        </table>
      </div>
    )
  }
}

export default Table
