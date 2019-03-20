import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import MenuItem from "./menuItem";
import { Theme } from '@material-ui/core/styles/createMuiTheme';

const styles = (theme:Theme) => ({
  container: {
    display: 'grid',
    gridTemplateColumns: 'repeat(4, 1fr)',
    gridGap: `${theme.spacing.unit * 3}px`,
  }
});

function Menu(props:any) {
  const { classes, data } = props;

  return (
    <div className={classes.container}>
      {data.map((cellData:any) =>
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
