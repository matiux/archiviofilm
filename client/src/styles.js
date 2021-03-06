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
      width: '80%',
      float: 'left',
      marginRight: '30px'
   },
   unseenCheckBox: {
      float: 'left',
      width: '10%'
   },
   unseenCheck: {
      margin: '0 auto',
   },
   item: {
      float: 'left'
   },
   switch: {
      root: {
         width: '50px',
         float: 'right',
         backgroundColor: 'red',
      }
   },
   treeStyle: {
      tree: {
         base: {
            listStyle: 'none',
            backgroundColor: '',
            margin: 0,
            padding: 0,
            color: '#757b7f',
            fontFamily: 'lucida grande ,tahoma,verdana,arial,sans-serif',
            fontSize: '14px'
         },
         node: {
            base: {
               //position: 'relative',
               //margin: '5px 0px 5px 0px',
               //padding: '4px',
            },
            link: {
               cursor: 'pointer',
               position: 'relative',
               padding: '4px',
               display: 'block',
               margin: '5px 0px 5px 0px',
               borderRadius: '5px',

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
                  fill: '#757b7f',
                  strokeWidth: 0
               }
            },
            header: {
               base: {
                  display: 'inline-block',
                  verticalAlign: 'top',
                  color: '#757b7f',
                  width: '100%'
               },
               baseTitle: {
                  display: 'inline-block',
                  verticalAlign: 'top',
                  color: '#757b7f',
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
