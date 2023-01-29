# Queer Global Configuration Library

This library provides an interface for accessing configuration values for Queer Global golang applications. 

The purpose of this library is to provide a uniform interface for reading config values, regardless of their source. 

Current versions of this library support reading values from config files (json) and from environment variables. Future work will likely include yaml support, as well as cloud-based config stores (s3 files, aws secrets and the like.)

## Usage

To use the config library, one must provide an init file which describes the type of configuration source, and the variables that should be collected. 

The general structure of an init file is as follows: 

```json
{
    "ConfigType" : "type-of-configuration-reader",
    "Environment" : "dev",
    "ConfigTTL" : "10s",
    "Aliases" : {
        "test1": "some.variable.name.or.path.test1", 
        "test2": "some.variable.name.or.path.test2"
    },
	"InitValues" : {
        "SomeKey" : "some configuration-type specific value"
    }
}
```

Where: 

- ConfigType describes the type of configuration (current allowed values are json or envvar)
- Environment is the environment for which this configuration applies.
- ConfigTTL is a go duration string which tells us how often the config library should refresh values from the source
- Aliases describe variable names we'd like to use to access a given config value from inside our program
- InitValues is an object containing values specific to the specific ConfigType. For example, json config would contain the file name from which to read. A hypothetical S3 bucket reader might contain the bucket name and config path.

### Examples

Example JSON-config init file:


```json
{
    "ConfigType" : "json",
    "Environment" : "unittest",
    "ConfigTTL" : "10s",
    "Aliases" : {
        "test1": "testvalues.test1", 
        "test2": "testvalues.testObj.test2"
    },
    "InitValues" : {
        "Path" : "unittest-config.json"
    }
}
```

The json init file described above instructs the library to read configuration from a json file called unittest-config.json

The config file itself might look like this:

```json
{
    "testvalues" : {
        "test1" : "test value 1",
        "testObj" : {
            "test2" : 2
        }
    },
    "test3" : "test value 3",
    "test4" : 4.0,
    "test5" : 5
}
```


An example EnvVar init file, which would instruct the library to read environment variables TEST_ENV_VAR and TEST_ENV_VAR2, and store aliases at keys test1 and test2.:

```json
{
	"ConfigType" : "envvar",
    "Environment" : "unittest",
    "ConfigTTL" : "10s",
    "Aliases" : {
        "test1": "TEST_ENV_VAR", 
        "test2": "TEST_ENV_VAR2"
    }
}
```
# Queer Global Configuration Library

This library provides an interface for accessing configuration values for Queer Global golang applications. 

The purpose of this library is to provide a uniform interface for reading config values, regardless of their source. 

Current versions of this library support reading values from config files (json) and from environment variables. Future work will likely include yaml support, as well as cloud-based config stores (s3 files, aws secrets and the like.)

## Usage

To use the config library, one must provide an init file which describes the type of configuration source, and the variables that should be collected. 

The general structure of an init file is as follows: 

```json
{
    "ConfigType" : "type-of-configuration-reader",
    "Environment" : "dev",
    "ConfigTTL" : "10s",
    "Aliases" : {
        "test1": "some.variable.name.or.path.test1", 
        "test2": "some.variable.name.or.path.test2"
    },
    "InitValues" : {
        "SomeKey" : "some configuration-type specific value"
    }
}
```

Where: 

- ConfigType describes the type of configuration (current allowed values are json or envvar)
-  Environment is the environment for which this configuration applies.
- ConfigTTL is a go duration string which tells us how often the config library should refresh values from the source
- Aliases describe variable names we'd like to use to access a given config value from inside our program
-  InitValues is an object containing values specific to the specific ConfigType. For example, json config would contain the file name from which to read. A hypothetical S3 bucket reader might contain the bucket name and config path.

### Examples

Example JSON-config init file:


```json
{
    "ConfigType" : "json",
    "Environment" : "unittest",
    "ConfigTTL" : "10s",
    "Aliases" : {
        "test1": "testvalues.test1", 
        "test2": "testvalues.testObj.test2"
    },
    "InitValues" : {
        "Path" : "unittest-config.json"
    }
}
```

The json init file described above instructs the library to read configuration from a json file called unittest-config.json

The config file itself might look like this:

```json
{
     "testvalues" : {
        "test1" : "test value 1",
        "testObj" : {
            "test2" : 2
      }
    },
    "test3" : "test value 3",
    "test4" : 4.0,
    "test5" : 5
}
```


An example EnvVar init file, which would instruct the library to read environment variables TEST_ENV_VAR and TEST_ENV_VAR2, and store aliases at keys test1 and test2.:

```json
{
    "ConfigType" : "envvar",
    "Environment" : "unittest",
    "ConfigTTL" : "10s",
    "Aliases" : {
        "test1": "TEST_ENV_VAR", 
        "test2": "TEST_ENV_VAR2"
    }
}
```

