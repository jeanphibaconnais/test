import React, { Component } from 'react';
import TodoListComplexForm from './TodoListComplexForm/TodoListComplexForm';
import TodoListComplexItems from './TodoListComplexItems/TodoListComplexItems';

class TodoListComplex extends Component {

    constructor(props) {
        super(props);
        this.state = {
            items : []
        };
    }
    addTodo = (newTodo) => {
        this.state.items.push(newTodo);
        this.setState({items: this.state.items});
    }

    render() {
        
        return <div className="todoListcomplex">
            <h3>Todo list multi components with data exchange between children to parent</h3>
            <TodoListComplexForm addTodoParent={this.addTodo}></TodoListComplexForm>        
            <TodoListComplexItems items={this.state.items}></TodoListComplexItems>
        </div>;
    }
}
export default TodoListComplex;