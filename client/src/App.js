import React, { Component } from 'react';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import { Toggle } from 'material-ui'
// import client, { updateFilm, fetchList } from "./client";
import { updateFilm, fetchList } from "./client";
import { Treebeard, decorators } from 'react-treebeard';
import styles from './styles';

import getMuiTheme from 'material-ui/styles/getMuiTheme';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';

function foo(list, search) {

   if (list.name === search) {

      return list;

   } else {

      if (list.children) {

         for (var i = 0; i < list.children.length; i++) {

            var l = foo(list.children[i], search);

            if (l) {
               return l;
            }
         }
      }
   }

   return false;
};

decorators.Toggle = () => (<span />);
decorators.Header = (props) => {

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
};

//decorators.Container = (props) => (<div><decorators.Toggle props={props} /><decorators.Header props={props} /></div>);

class FilmTree extends Component {

   constructor(props) {

      super(props);

      this.state = {
         films: {},
      };
   }

   componentDidMount() {

      fetchList()

         .then((response) => {

            var hierarchy = response.data.reduce(function (hier, film) {

               var x = hier;
               var path = film.File.replace('/mnt/storaMioArchivio/MieiVideo/Film/', '');
               var elements = path.split('/');

               var parts = elements.length;

               elements.forEach(function (item, i) {

                  var n = foo(x, item);

                  if (item !== n.name) {

                     var l = {
                        name: item,
                        id: film.Id,
                        seen: film.Seen,
                        path: film.File
                     };

                     if (i + 1 !== parts) {

                        l['children'] = [];
                     }

                     n ? n.children.push(l) : x.children.push(l);

                     x = l;

                  } else {

                     x = n;
                  }
               });

               return hier;

            }, { name: '/', children: [], toggled: true, });

            this.setState({ films: hierarchy })

            console.log(this.state);

            // this.setState((prevState, props) => {

            //   return { films: response.data };
            // });

         })
         .catch(function (error) {

            console.log(error);
         });
   }

   onToggle = (node, toggled) => {

      if (this.state.cursor) {

         this.setState({ cursor: { active: false } })
      }

      //node.active = true;

      if (node.children) {

         node.toggled = toggled;
      }

      this.setState({ cursor: node });
   };

   render() {

      return (
         <div>
            <div style={styles.searchBox}>
               <div className="input-group">
                  <span className="input-group-addon">
                     <i className="fa fa-search"></i>
                  </span>
                  <input type="text"
                     className="form-control"
                     placeholder="Search the tree..."
                  //onKeyUp={this.onFilterMouseUp.bind(this)}
                  />
               </div>
            </div>
            <div style={styles.component}>
               <Treebeard
                  data={this.state.films}
                  onToggle={this.onToggle}
                  decorators={decorators}
                  style={styles.treeStyle}
               />
            </div>
         </div>
      );
   }
}



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

class App extends Component {

   render() {

      return (
         <MuiThemeProvider muiTheme={getMuiTheme(darkBaseTheme)}>
            <FilmTree />
         </MuiThemeProvider>
      );
   }
}

export default App;