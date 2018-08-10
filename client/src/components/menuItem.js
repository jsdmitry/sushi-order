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
import TextField from '@material-ui/core/TextField';

const styles = theme => ({
  card: {
    maxWidth: 300
  },
  media: {
    height: 0,
    paddingTop: '56.25%', // 16:9
  },
  input: {
    width: 40
  },
  footer: {
    width: "100%",
    display: 'grid',
    gridTemplateColumns: '2fr 1fr 1fr',
  },
  count: {
    marginTop: "-16px"
  },
  cart: {
    display: "flex",
    justifyContent: "flex-end",
    marginTop: "-4px"
  },
  price: {
    marginTop: 8,
    marginLeft: 8
  }
});

function MenuItem(props) {
  const { classes, imageURL, title, description, price } = props;
  return (
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
        </CardContent>
        <CardActions>
          <div className={classes.footer}>
            <Typography
              variant="subheading"
              align="left"
              className={classes.price}
            >
                {price + "руб."}
            </Typography>
            <TextField
              className={classes.count}
              id="number"
              label="Кол-во"
              //value={this.state.age}
              //onChange={this.handleChange('age')}
              type="number"
              //className={classes.textField}
              margin="none"
            />
            <div className={classes.cart}>
              <IconButton aria-label="Cart">
                <ShoppingCartOutlined />
              </IconButton>
            </div>
          </div>
        </CardActions>
      </Card>
  );
}

MenuItem.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(MenuItem);
