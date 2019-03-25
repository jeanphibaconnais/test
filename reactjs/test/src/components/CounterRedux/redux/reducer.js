import { INCREMENT, DECREMENT } from './actions';

const initialState = {
  count: 0
};

function reducer(state = initialState, action) {
  switch(action.type) {
    case INCREMENT.type:
      return {
        count: state.count + 1
      };
    case DECREMENT.type:
      return {
        count: state.count - 1
      };
    default:
      return state;
  }
}

export default reducer;