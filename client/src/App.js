import React, { Component } from 'react';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import { LabelCheckbox } from 'material-ui/Checkbox'
import Switch from './CustomSwitch'
// import client, { updateFilm, fetchList } from "./client";
import { updateFilm, fetchList } from "./client";
import { Treebeard, decorators } from 'react-treebeard';
import styles from './styles';
import { debounce } from 'throttle-debounce';

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

   const style = props.style.header;
   const iconType = props.node.children ? 'folder' : 'file-text';
   const iconClass = `fa fa-${iconType}`;
   const iconStyle = { marginRight: '5px' };

   const film = {
      seen: props.node.seen,
      id: props.node.id,
      path: props.node.path
   };

   //console.log(style)

   return (
      <div className="bar" style={iconType === 'folder' ? style.baseTitle : style.base}>
         <div className="baz" style={style.title}>
            <div className="boo" style={styles.item}>
               <i className={iconClass} style={iconStyle} />
               {props.node.name}
            </div>
            {!props.node.children ? <FilmSeenStatusToggle film={film} /> : ''}
         </div>
      </div>
   );
};

decorators.Container = (props) => {

   return (

      <div className={props.node.children ? "folderElement" : "childrenElement"} style={props.style.link} onClick={props.onClick}>
         <decorators.Toggle className="ecc" {...props} />
         <decorators.Header className="laa" {...props} />
      </div>

   );
}

class FilmTree extends Component {

   constructor(props) {

      super(props);

      this.filterByName = debounce(1000, this.filterByName);

      this.state = {
         films: {},
         unseen: false,
         filter: null,
      };
   }

   componentDidUpdate(prevProps, prevState) {

      if (prevState.unseen !== this.state.unseen) {

         this.fetchData()
      }
   }

   componentDidMount() {

      this.fetchData()
   }

   fetchData() {

      fetchList(this.state.unseen, this.state.filter)

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

            //console.log(this.state);

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

   toggleUnseen = (event) => {

      const check = event.target.checked;

      this.setState({ unseen: check });
   }

   debounceFilterByName = (e) => {

      this.filterByName(e.target.value);
   }

   filterByName = (filter) => {

      var f = filter.trim(filter);

      this.setState({ filter: f ? f : null });

      this.fetchData()
   }

   render() {

      return (
         <div>
            <div style={styles.filters}>
               <div style={styles.searchBox}>
                  <div className="input-group">
                     <span className="input-group-addon">
                        <i className="fa fa-search"></i>
                     </span>
                     <input
                        type="text"
                        className="form-control"
                        placeholder="Search the tree..."
                        onKeyUp={this.debounceFilterByName}
                     />
                  </div>
               </div>
               <div style={styles.unseenCheckBox}>
                  <LabelCheckbox
                     onChange={this.toggleUnseen}
                     label="Unseen"
                     checked={this.state.unseen}
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
         <Switch onChange={this.changeSeenStatus} {...this.state} classNane={styles.switch.root}>
         </Switch>
      );
   }
}

class App extends Component {

   render() {

      return (
         <MuiThemeProvider>
            <FilmTree />
         </MuiThemeProvider>
      );
   }
}

export default App;