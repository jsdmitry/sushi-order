import React from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import {withStyles} from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import { Theme } from '@material-ui/core/styles/createMuiTheme';

const styles = (theme:Theme) => ({
  root: {
    width: '100%',
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,
  },
  avatar: {
    width: 60,
    height: 60
  }
});

function CategoriesList(props:any) {
  const { classes, categories, onCategorySelectionChanged } = props;
  const listItems = categories.data.map((category:any) =>
    <ListItem
      button
      onClick={() => {
        onCategorySelectionChanged(category.ID);
      }}
      key={category.ID}
    >
      <Avatar
        src={category.ImageURL}
        className={classNames(classes.avatar)}
      >
      </Avatar>
      <ListItemText primary={category.Caption}/>
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
