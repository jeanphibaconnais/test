import React from 'react';

const TodoListComplexItems = (props) => {
    
    var items = props.items;

    return (
        <ul>
            {items.map((value) => {
                return <li>{value}</li>
            })} 
        </ul>
    );

}
export default (TodoListComplexItems);