export default {
   component: {
      width: '100%',
      display: 'inline-block',
      verticalAlign: 'top',
      padding: '20px',
      '@media (maxWidth: 640px)': {
         width: '100%',
         display: 'block'
      }
   },
   filters: {

      padding: '20px 20px 0 20px',
      width: '100%'
   },
   searchBox: {
      width: '90%',
      float: 'left',
   },
   unseenCheckBox: {
      float: 'left',
      width: '10%'
   },
   unseenCheck: {
      //border: '2px solid #FF9800',
      margin: '0 auto',
   },
   item: {
      float: 'left'
   },
   toggle: {
      width: 'auto',
      float: 'right'
   },
   treeStyle: {
      tree: {
         base: {
            listStyle: 'none',
            backgroundColor: '',
            margin: 0,
            padding: 0,
            color: '#9DA5AB',
            fontFamily: 'lucida grande ,tahoma,verdana,arial,sans-serif',
            fontSize: '14px'
         },
         node: {
            base: {
               position: 'relative',
               borderWidth: '1px',
               border: '1px black solid',
               margin: '5px 0px 5px 0px',
               padding: '4px',
            },
            link: {
               cursor: 'pointer',
               position: 'relative',
               padding: '0px 5px',
               display: 'block',

            },
            activeLink: {
               background: '#31363F'
            },
            toggle: {
               base: {
                  position: 'relative',
                  display: 'inline-block',
                  verticalAlign: 'top',
                  marginLeft: '-5px',
                  height: '24px',
                  width: '24px'
               },
               wrapper: {
                  position: 'absolute',
                  top: '50%',
                  left: '50%',
                  margin: '-7px 0 0 -7px',
                  height: '14px'
               },
               height: 14,
               width: 14,
               arrow: {
                  fill: '#9DA5AB',
                  strokeWidth: 0
               }
            },
            header: {
               base: {
                  display: 'inline-block',
                  verticalAlign: 'top',
                  color: '#9DA5AB',
                  width: '100%'
               },
               baseTitle: {
                  display: 'inline-block',
                  verticalAlign: 'top',
                  color: '#9DA5AB',
                  width: '100%',
                  fontWeight: 'bold',
               },
               connector: {
                  width: '2px',
                  height: '12px',
                  borderLeft: 'solid 2px black',
                  borderBottom: 'solid 2px black',
                  position: 'absolute',
                  top: '0px',
                  left: '-21px'
               },
               title: {
                  lineHeight: '24px',
                  verticalAlign: 'middle',

               }
            },
            subtree: {
               listStyle: 'none',
               paddingLeft: '19px'
            },
            loading: {
               color: '#E2C089'
            }
         }
      }
   }
};
