import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import MenuItemRow from '../components/menuItemRow';
import Grid from '@material-ui/core/Grid';

const styles = theme => ({
  root: {
    flexGrow: 1
  }
});

function prepareMenuData(data, colCount) {
  let result = [];
  for(let i = 0; i < data.length; i += colCount) {
    result.push(data.slice(i,  i + colCount));
  }
  return result;
}

function Menu(props) {
  const { classes, data, colCount } = props;
  const rowsData = prepareMenuData(data, colCount);

  return (
    <div className={classes.root}>
      <Grid container spacing={8}>
      {rowsData.map((rowData, index) =>
        <MenuItemRow data={rowData} spacing={8} key={index}/>
      )}
      </Grid>
    </div>
  );
}

Menu.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Menu);
