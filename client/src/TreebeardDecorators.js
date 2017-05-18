import React, { Component, PropTypes } from 'react';
import { TreeBeard } from 'react-treebeard';
import _ from 'lodash';
import styles from './styles';
import { updateFilm } from "./client";

class FilmSeenStatusToggle extends Component {

   constructor(props) {

      super(props);

      this.state = {
         seen: props.film.seen,
         id: props.film.id,
         path: props.film.path
      };
   }

   changeSeenStatus = (event) => {

      const check = event.target.checked;

      this.setState({ seen: check });

      updateFilm(this.state.id, { "Seen": check })
   }

   render() {

      return (

         <Toggle onToggle={this.changeSeenStatus} defaultToggled={this.state.seen} style={styles.toggle} />

      );
   }
}

const Loading = (props) => {
   return (
      <div style={props.style}>
         loading...
        </div>
   );
};

Loading.displayName = 'Loading';

Loading.propTypes = {
   style: PropTypes.object
};

const Toggle = (props) => {

   return (<span />)
}

Toggle.displayName = 'Toggle';
Toggle.propTypes = {
   node: PropTypes.object,
   style: PropTypes.object
};

function Header(props) {
   const style = props.style;
   const iconType = props.node.children ? 'folder' : 'file-text';
   const iconClass = `fa fa-${iconType}`;
   const iconStyle = { marginRight: '5px' };

   const film = {
      seen: props.node.seen,
      id: props.node.id,
      path: props.node.path
   };

   //var styleBase = Object.clone(style.base)

   //var s = props.node.children ? Object.assign(style.base, {backgroundColor: '#b3f442'}) : style.base;

   return (
      <div style={iconType === 'folder' ? props.style.baseTitle : props.style.base}>
         <div style={style.title}>
            <div style={styles.item}>
               <i className={iconClass} style={iconStyle} />
               {props.node.name}
            </div>
            {!props.node.children ? <FilmSeenStatusToggle film={film} /> : ''}
         </div>
      </div>
   );
}

Header.displayName = 'TreeHeader';
Header.propTypes = {
   node: PropTypes.object.isRequired
};


class Container extends Component {
   renderToggle() {
      return this.renderToggleDecorator();
   }

   renderToggleDecorator() {
      const { style, decorators, node } = this.props;
      return (<decorators.Toggle node={node} style={style.toggle} />);
   }

   render() {
      const { style, decorators, terminal, onClick, node } = this.props;
      const finalStyle = style.container.reduce((total, rules) => {
         return _.assign(total, rules);
      }
         , {});

      return (
         <div className='tree-container'>
            <span ref="clickable"
               onClick={onClick}
               style={finalStyle}>
               {!terminal ? this.renderToggle() : null}
            </span>
            <decorators.Header
               node={node}
               style={style.header}
            />
         </div>
      );
   }
}

Container.propTypes = {
   decorators: PropTypes.object,
   node: PropTypes.object,
   onClick: PropTypes.func,
   style: PropTypes.object,
   terminal: PropTypes.bool
};

const decorators = {
   Container,
   Header,
   Loading,
   Toggle
};

export default class CustomTree extends Component {
   onToggle = (node, toggled) => {
      this.props.onToggle({ name: node.slug, toggled });
   };

   render() {

      console.log(this.props);

      return (
         <div className="mailboxes-list">
            <TreeBeard
               data={this.props.data}
               onToggle={this.onToggle}
               decorators={decorators}
            />
         </div>
      );
   }
}

CustomTree.propTypes = {
   data: PropTypes.object,
   onToggle: PropTypes.func
};