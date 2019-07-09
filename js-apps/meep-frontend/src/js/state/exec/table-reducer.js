/*
 * Copyright (c) 2019
 * InterDigital Communications, Inc.
 * All rights reserved.
 *
 * The information provided herein is the proprietary and confidential
 * information of InterDigital Communications, Inc.
 */
import { updateObject } from '../../util/object-util';

const initialState = {
  data: [],
  selected: [],
  order: 'asc',
  orderBy: 'name',
  rowsPerPage: 10,
  page: 0
};

const EXEC_CHANGE_TABLE = 'EXEC_CHANGE_TABLE';
function execChangeTable(table) {
  return {
    type: EXEC_CHANGE_TABLE,
    payload: table
  };
}

export { execChangeTable, EXEC_CHANGE_TABLE };

export function execTableReducer(state = initialState, action) {
  switch (action.type) {
  case EXEC_CHANGE_TABLE:
    return updateObject({}, action.payload);
  default:
    return state;
  }
}