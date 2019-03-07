import React, { Component } from 'react';

class ListTerm extends Component {
    
    render() {
        let list;

        list = this.props.items.map((item) => {
            return <li key={item}>{item}</li>;
        });

        return <div>{list}</div>;
    }
}

export default ListTerm;
