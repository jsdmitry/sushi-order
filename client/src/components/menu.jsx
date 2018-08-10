import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import MenuItem from "./menuItem";

const styles = theme => ({
  container: {
    display: 'grid',
    gridTemplateColumns: 'repeat(4, 1fr)',
    gridGap: `${theme.spacing.unit * 3}px`,
  }
});

function Menu(props) {
  const { classes, data } = props;

  return (
    <div className={classes.container}>
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
    </div>
  );
}

Menu.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Menu);
