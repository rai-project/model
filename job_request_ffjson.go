// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: job_request.go

package model

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *BuildImageSpecification) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *BuildImageSpecification) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"image_name":`)
	fflib.WriteJsonString(buf, string(j.ImageName))
	buf.WriteString(`,"dockerfile":`)
	fflib.WriteJsonString(buf, string(j.Dockerfile))
	if j.Push != nil {
		buf.WriteString(`,"push":`)

		{

			err = j.Push.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"push":null`)
	}
	if j.NoCache {
		buf.WriteString(`,"no_cache":true`)
	} else {
		buf.WriteString(`,"no_cache":false`)
	}
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *BuildSpecification) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *BuildSpecification) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"rai":`)

	{

		err = j.RAI.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"resources":`)

	{

		err = j.Resources.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"commands":`)

	{

		err = j.Commands.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *CPUResources) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *CPUResources) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"architecture":`)
	fflib.WriteJsonString(buf, string(j.Architecture))
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *CommandsBuildSpecification) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *CommandsBuildSpecification) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if j.BuildImage != nil {
		if true {
			buf.WriteString(`"build_image":`)

			{

				err = j.BuildImage.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.WriteString(`"build":`)
	if j.Build != nil {
		buf.WriteString(`[`)
		for i, v := range j.Build {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *GPUResources) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *GPUResources) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"architecture":`)
	fflib.WriteJsonString(buf, string(j.Architecture))
	buf.WriteString(`,"count":`)
	fflib.FormatBits2(buf, uint64(j.Count), 10, j.Count < 0)
	buf.WriteString(`,"memory":`)
	fflib.FormatBits2(buf, uint64(j.Memory), 10, j.Memory < 0)
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *JobRequest) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *JobRequest) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	/* Struct fall back. type=config.VersionInfo kind=struct */
	buf.WriteString(`{"client_version":`)
	err = buf.Encode(&j.ClientVersion)
	if err != nil {
		return err
	}
	buf.WriteString(`,"upload_key":`)
	fflib.WriteJsonString(buf, string(j.UploadKey))
	buf.WriteString(`,"user":`)

	{

		err = j.User.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"build_specification":`)

	{

		err = j.BuildSpecification.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"id":`)
	fflib.WriteJsonString(buf, string(j.ID))
	buf.WriteString(`,"created_at":`)

	{

		obj, err = j.CreatedAt.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"updated_at":`)

	{

		obj, err = j.UpdatedAt.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	if j.DeletedAt != nil {
		buf.WriteString(`,"deleted_at":`)

		{

			obj, err = j.DeletedAt.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
	} else {
		buf.WriteString(`,"deleted_at":null`)
	}
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *Push) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Push) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	if j.Push {
		buf.WriteString(`{"push":true`)
	} else {
		buf.WriteString(`{"push":false`)
	}
	buf.WriteString(`,"image_name":`)
	fflib.WriteJsonString(buf, string(j.ImageName))
	buf.WriteString(`,"registry":`)
	fflib.WriteJsonString(buf, string(j.Registry))
	buf.WriteString(`,"credentials":`)

	{

		err = j.Credentials.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *RAIBuildSpecification) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *RAIBuildSpecification) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"version":`)
	fflib.WriteJsonString(buf, string(j.Version))
	buf.WriteString(`,"image":`)
	fflib.WriteJsonString(buf, string(j.ContainerImage))
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *Resources) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Resources) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"cpu":`)

	{

		err = j.CPU.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	if j.GPU != nil {
		buf.WriteString(`,"gpu":`)

		{

			err = j.GPU.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"gpu":null`)
	}
	if j.Network {
		buf.WriteString(`,"network":true`)
	} else {
		buf.WriteString(`,"network":false`)
	}
	buf.WriteByte('}')
	return nil
}
