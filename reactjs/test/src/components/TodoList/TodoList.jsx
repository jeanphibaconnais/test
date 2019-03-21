import React, { Component } from 'react';
import TodoListItems from './TodoListItems/TodoListItems';

class TodoList extends Component {

    constructor(props) {
        super(props);
        this.state = {
            items : [],
            currentElement: ''
        };
    }

    handleChange = (event) => {
        this.setState({currentElement : event.target.value});
    }

    onSubmit = (event) => {
        //event.eventPreventDefault(); // permet de bloquer le fonctionnement normal du formulaire
        this.setState({
            currentElement: '',
            items: [ ...this.state.items, this.state.currentElement ]
        });
    };

    render() {
        
        return <div className="todoList">
                <h3>TodoList</h3>

                <span>Current : {this.state.currentElement}</span>
                <br/>
                
                <input type="text" value={this.state.currentElement} onChange={this.handleChange} />
                <button type="submit" onClick={this.onSubmit}>Add element</button>

                <TodoListItems items={this.state.items}></TodoListItems>
            </div>
        ;
    }
}

export default TodoList;