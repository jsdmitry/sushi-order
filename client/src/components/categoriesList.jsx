import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';

const styles = theme => ({
  root: {
    width: '100%',
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,
  }
});

function CategoriesList(props) {
  const { classes, categories } = props;
  const listItems = categories.data.map((category) =>
    <ListItem button>
      <Avatar src={category.ImageURL}>
      </Avatar>
      <ListItemText primary={category.Caption} />
    </ListItem>
  );
  return (
    <div className={classes.root}>
      <List component="nav">
        {listItems}
      </List>
    </div>
  );
}

CategoriesList.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(CategoriesList);
