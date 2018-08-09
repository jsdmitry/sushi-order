import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import MenuItem from '../components/menuItem';

const styles = theme => ({
  root: {
    flexGrow: 1,
  }
});

function MenuItemRow(props) {
  const { classes, data, spacing } = props;

  return (
    <Grid
      className={classes.root}
      container
      spacing={spacing}
      alignItems="center"
      direction="row"
      justify="center"
    >
      {data.map((cellData) =>
        <MenuItem
          title={cellData.Caption}
          imageURL={cellData.ImageURL}
          description={cellData.Description}
          price={cellData.Price}
          key={cellData.Caption}
        >
        </MenuItem>
      )}
    </Grid>

  );
}

MenuItemRow.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(MenuItemRow);
