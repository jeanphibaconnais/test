import React, { Component } from 'react';
import ListTerms from '../ListsTerm/ListTerm';

export default class HelloWorld extends Component {
    constructor(props) {
        super(props);
        this.state = {
            online: true,
            counter: 0,
            term: 'coucou',
            items: []
        };
    }

    callbackclick() {
        console.log('erreur ...');
    }
    
    handleChange(event) {
        
    }

    onSubmit = (event) => {
        //event.eventPreventDefault(); // permet de bloquer le fonctionnement normal du formulaire
        this.setState({
            term: '',
            items: [ ...this.state.items, this.state.term ]
        });
    };

    render() {
        const user = this.props.user;
        let classNameSpan = 'span';

        return (
            <div>
                <p>
                    {' '}
                    Component with props : hello {user}, you are {!this.state.online || 'online'} !{' '}
                </p>
                <span className={classNameSpan}>{this.state.counter}</span>
                
                <button onClick={() => { this.setState({ counter: this.state.counter + 1 }); }}>Increment this one
                </button>

                <form onSubmit={this.onSubmit}>
                    <input value={this.state.term} onChange={this.handleChange} />
                    <input type="submit" value="Ajouter" />
                </form>

                <ListTerms items={this.state.items} />
            </div>
        );
    }
}
