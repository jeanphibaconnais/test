### Premiers tests Reactjs


#### Pages m'ayant aidé à appréhender Reactjs

- https://blogs.infinitesquare.com/posts/web/premiers-pas-avec-react-js-partie-2

Schéma du cycle de vie : https://infiniteblogs.blob.core.windows.net/medias/ccff968e-b694-45fa-b988-f5faa8966c9d_LIFE%20CYCLE.png

- https://putaindecode.io/fr/articles/js/react/

- https://blog.ippon.fr/2016/04/19/comment-bien-debuter-en-reactjs/

- https://openclassrooms.com/fr/courses/4664381-realisez-une-application-web-avec-react-js/4664831-reagissez-aux-evenements

### Création de composant 
en classes
en function

#### Props Les states

Les props permettent de faire transiter des informations des éléments parents vers les composants enfants.
Et pas l'inverse !  Les props sont immutables.

Le state permet de stocker des informations d'un composant et de les modifier. Le state est donc mutable. Pour modifier un state, this.setState(data, callback).


#### Hooks intéressant 
Dans les composants créés avec des classes, des hooks sont disponibles. Ils permettent de suivre le cycle dev vie des composants Reacts.

ComponentWillReceiveProps : Peut éventuellement servir à vérifier les données saisies à l'écran avant d'enchainer sur d'autres écrans.

ComponentDidUpdate : peut permettre de mettre à jour les nouvelles props envoyer par le composant parent.

ComponentWillUnmount : permet de nettoyer les éléments avant la suppression du composant.

### Routing
TODO 

### Hooks : nouveauté majeure React avec la v16.8
https://reactjs.org/docs/hooks-intro.html

Permettre d'écrire des hooks sans avoir de classe
 => fonction useState

Facebook déconseille de récrire vos composants en function !

 => function useEffect : Similar to componentDidMount and componentDidUpdate 
 Généralement les actions à faire à l'init du composant sont semblables à celles faites lors de l'update.  

 Permet de créer ses propres hook. Cela permet d'isoler des fonctionnalités comportementales et de les tester plus facilement.

 => function useReducer : permet de gérer son local storage

 Passer par des functions, a un avantage de lisibilté. 
 Par exemple ce code :
```
 const [count, setCount] = useState(0);
```

 Dans des classes l'accès à une variable se fait via {{this.state.mavar}}. Dans des functions, cela devient {{mavar}}.

Mise à jour du state plus rapide : this.setState({count : count + 1}) remplacé par setCount(count + 1)




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
