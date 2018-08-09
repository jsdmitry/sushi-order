import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import IconButton from '@material-ui/core/IconButton';
import ShoppingCartOutlined from '@material-ui/icons/ShoppingCartOutlined';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';

const styles = {
  card: {
    maxWidth: 300,
  },
  media: {
    height: 0,
    paddingTop: '56.25%', // 16:9
  }
};

function MenuItem(props) {
  const { classes, imageURL, title, description, price } = props;
  return (
    <Grid item xs>
      <Card className={classes.card}>
        <CardMedia
          className={classes.media}
          image={imageURL}
          title={title}
        />
        <CardContent>
          <Typography gutterBottom variant="headline" component="h2">
            {title}
          </Typography>
          <Typography component="p">
            {description}
          </Typography>
          <Typography component="h1">
            {price} руб.
          </Typography>
        </CardContent>
        <CardActions>
          <IconButton aria-label="Cart">
            <ShoppingCartOutlined />
          </IconButton>
        </CardActions>
      </Card>
    </Grid>
  );
}

MenuItem.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(MenuItem);
