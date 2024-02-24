# Latipe Delivery Service 


> ### Tech Stack:
> - Programming Language: Go 1.20
> - Web Framework: Fiber v2
> - gRPC: v1.62
> - Database: MongoDB
> - Config: Viper
> - Tools: RabbitMQ, Prometheus, Grafana
> - Environment: Docker

### Features:
Managing the delivery function in e-commerce using a microservices architecture, including:
- **Providing Vietnam Location Data:** Offering provinces, districts, and wards data in Vietnam.
- **Shipping Cost:** Calculating shipping costs from store/vendor to buyer.
- **Delivery Management:** Performing CRUD operations on delivery entities.
### API Documentation:
- **Base URL:** http://localhost:5005
- **Swagger:** [API Documentation](http://localhost:5005/swagger/index.html)
- **API Endpoints:**
    - **Shipping Cost:**
        - `GET /api/v1/shipping-cost?from={from}&to={to}&weight={weight}`: Get shipping cost from store/vendor to buyer.
    - **Vietnam Location Data:**
        - `GET /api/v1/vietnam-location/provinces`: Get all provinces in Vietnam.
        - `GET /api/v1/vietnam-location/districts?province={province}`: Get all districts in a province.
        - `GET /api/v1/vietnam-location/wards?district={district}`: Get all wards in a district.
    - **Delivery Management:**
        - `GET /api/v1/delivery`: Get all delivery entities.
        - `GET /api/v1/delivery/{id}`: Get a delivery entity by id.
        - `POST /api/v1/delivery`: Create a new delivery entity.
        - `PUT /api/v1/delivery/{id}`: Update a delivery entity by id.
        - `DELETE /api/v1/delivery/{id}`: Delete a delivery entity by id.
    - **Metrics:**
        - `GET /metrics`: Get metrics data for monitoring the application, you can also use prometheus.
        - `GET /health`: Get health check data for monitoring the application.
        - `GET /readiness`: Get ready check.
        - `GET /liveness`: Get live check.
        - `GET /fiber/dashboard`: Get fiber dashboard for monitoring the application.
### Setup:
Change your configuration in the `config.yaml` file. Then, run the following command in make file:
```bash 
    make setup #setup the environment and get dependencies
    make buildw #build the application run on windows (use make bulidl for linux)
    make runw #run the application or run the binary file in /build folder
```
You can also use docker file to build and run the application:
```bash
    docker build -t latipe-delivery-service .
    docker run -p 5005:5005 latipe-delivery-service
```
<hr></hr>
<h4> Application Developed by Tran Tien Dat </h4>
