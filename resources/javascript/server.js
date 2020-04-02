const express = require('express')
const YAML = require('yamljs')
const swaggerUi = require('swagger-ui-express')
const { connector } = require('swagger-routes-express')
const api = require('./api')
const bodyParser = require('body-parser');


const makeApp = () => {
    const apiDefinition = YAML.load('../spec/spec.yml');
    const connect = connector(api, apiDefinition);
    const app = express();
    app.use('/api/ui', swaggerUi.serve, swaggerUi.setup(apiDefinition));
    app.use(bodyParser.urlencoded({ extended: true }));
    app.use(bodyParser.json());
    connect(app);
    return app
};

makeApp().listen(3001, function () {
    console.log('Example app listening on port 3001!');
});