# Welcome to PowerDale

PowerDale is a small town with around 100 residents. Most houses have a smartmeter installed that can save and send information
about how much energy a house has used.

There are three major providers of energy in town that charge different amounts for the power they supply.

- _Dr Evil's Dark Energy_
- _The Green Eco_
- _Power for Everyone_

## Introducing JOI Energy

JOI Energy is a new startup in the energy industry.
Rather than selling energy they want to differentiate themselves from the market by recording their customers' energy usage from their smart meters and
recommending the best suppler to meet their needs.

You have been placed into their development team, whose current goal is to produce an API which their customers and smart meters will interact with.

Unfortunately, two of the team are on annual leave, and another has called in sick!
You are left with a ThoughtWorker to progress with the current user stories on the story wall. This is your chance to make an impact on the business, improve the code base and deliver value.

## Learn More

To learn more about this project, take a look at the following resources:

- [Local Install](INSTALL.md) - Learn about how to install and use this project.
- [Architecture](ARCHITECTURE.md) - See all details about what technologies were used to build this project.

## Users

To trial the new JOI software 5 people from the JOI accounts team have agreed to test the service and share their energy data.

- Sarah - Smart Meter Id: "smart-meter-0", current power supplier: Dr Evil's Dark Energy.
- Peter - Smart Meter Id: "smart-meter-1", current power supplier: The Green Eco.
- Charlie - Smart Meter Id: "smart-meter-2", current power supplier: Dr Evil's Dark Energy.
- Andrea - Smart Meter Id: "smart-meter-3", current power supplier: Power for Everyone.
- Alex - Smart Meter Id: "smart-meter-4", current power supplier: The Green Eco.

### Store Readings

#### Endpoint

```
POST
/readings/store
```

#### Input

```json
{
    "smartMeterId": <smartMeterId>,
    "electricityReadings": [
        { "time": <UTC-Time>, "reading": <reading> },
        ...
    ]
}
```

`time`: UTC time, e.g. `2024-04-13T19:41:26Z`
`reading`: kW reading of smart meter at that time, e.g. `0.0503`

## Versions

We use [Semantic version](http://semver.org) for versioning. For versions available, see [changelog](CHANGELOG.md).

## License

This project is licensed under the [Creative Commons License](LICENSE).
