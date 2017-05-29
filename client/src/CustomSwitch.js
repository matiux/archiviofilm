import React from 'react';
import PropTypes from 'prop-types';
import { withStyles, createStyleSheet } from 'material-ui/styles';
import Switch from 'material-ui/Switch';

const styleSheet = createStyleSheet('OverridesClasses', {
   root: {
      //  background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
      //  borderRadius: 3,
      //  border: 0,
      //  color: 'white',
      //  height: 48,
      //  padding: '0 30px',
      //  boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .30)',
      float: 'right',
      display: 'inline-flex',
      width: 62,
      position: 'relative'
   }
});

function OverridesClasses(props) {

   return (
      <Switch classes={props.classes} onChange={props.onChange} checked={props.seen}>
         {props.children ? props.children : 'class names'}
      </Switch>
   );
}

OverridesClasses.propTypes = {
   children: PropTypes.node,
   classes: PropTypes.object.isRequired,
};

export default withStyles(styleSheet)(OverridesClasses);