### Premiers tests Reactjs


#### Pages m'ayant aidé à appréhender Reactjs

- https://blogs.infinitesquare.com/posts/web/premiers-pas-avec-react-js-partie-2

Schéma du cycle de vie : https://infiniteblogs.blob.core.windows.net/medias/ccff968e-b694-45fa-b988-f5faa8966c9d_LIFE%20CYCLE.png

#### Props Les states

Les props permettent de faire transiter des informations des éléments parents vers les composants enfants.
Et pas l'inverse !  Les props sont immutables.

Le state permet de stocker des informations d'un composant et de les modifier. Le state est donc mutable. Pour modifier un state, this.setState(data, callback).

#### Hooks intéressant 

ComponentWillReceiveProps : Peut éventuellement servir à vérifier les données saisies à l'écran avant d'enchainer sur d'autres écrans.

ComponentDidUpdate : peut permettre de mettre à jour les nouvelles props envoyer par le composant parent.

ComponentWillUnmount : permet de nettoyer les éléments avant la suppression du composant.



#### Commandes de bases npm pour le projet

### `npm start`

Runs the app in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

### `npm run build`

Builds the app for production to the `build` folder.<br>
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br>
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.
