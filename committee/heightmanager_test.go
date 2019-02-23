package committee

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewHeightManager(t *testing.T) {
	require := require.New(t)
	t.Run("success-new-height-manager", func(t *testing.T) {
		hm := newHeightManager()
		require.NotNil(hm)
	})
}
func TestValidate(t *testing.T) {
	require := require.New(t)
	t.Run("success-to-validate", func(t *testing.T) {
		hm := newHeightManager()
		require.NotNil(hm)
		require.Nil(hm.validate(10, time.Now()))

	})
	t.Run("failed-to-validate", func(t *testing.T) {
		hm := newHeightManager()
		startTime := time.Now()
		require.NotNil(hm)
		t.Run("failed-new-invalid-time", func(t *testing.T) {
			require.NoError(hm.validate(2, startTime))
			require.Nil(hm.add(2, startTime))
			require.Error(hm.validate(3, startTime.Local().AddDate(0, 0, -1)))
		})
	})
	t.Run("failed-to-validate", func(t *testing.T) {
		hm := newHeightManager()
		startTime := time.Now()
		require.NotNil(hm)
		t.Run("failed-new-invalid-height", func(t *testing.T) {
			require.NoError(hm.validate(2, startTime))
			require.Nil(hm.add(2, startTime))
			require.Error(hm.validate(1, startTime.Local().AddDate(0, 0, 1)))
		})
	})
	t.Run("success-to-validate", func(t *testing.T) {
		hm := newHeightManager()
		startTime := time.Now()
		require.NotNil(hm)
		require.NoError(hm.validate(2, startTime))
		require.Nil(hm.add(2, startTime))
		require.NoError(hm.validate(3, startTime.Local().AddDate(0, 0, 1)))
		require.Nil(hm.add(3, startTime.Local().AddDate(0, 0, 1)))
		require.NoError(hm.validate(4, startTime.Local().AddDate(0, 0, 2)))
	})
}

func TestAdd(t *testing.T) {
	require := require.New(t)
	startTime := time.Now()
	t.Run("success-new-height-manager", func(t *testing.T) {
		hm := newHeightManager()
		require.NotNil(hm)
		require.Nil(hm.add(2, startTime))
		require.Nil(hm.add(3, startTime.Local().AddDate(0, 0, 1)))
		require.Nil(hm.add(4, startTime.Local().AddDate(0, 0, 2)))
		require.Nil(hm.add(5, startTime.Local().AddDate(0, 0, 3)))
	})
	t.Run("failed-new-invalid-height", func(t *testing.T) {
		hm := newHeightManager()
		require.NotNil(hm)
		require.Nil(hm.add(2, startTime))
		require.Nil(hm.add(4, startTime.Local().AddDate(0, 0, 1)))
		require.Error(hm.add(3, startTime.Local().AddDate(0, 0, 2)))
		require.Nil(hm.add(5, startTime.Local().AddDate(0, 0, 3)))
	})
	t.Run("failed-new-invalid-time", func(t *testing.T) {
		hm := newHeightManager()
		require.NotNil(hm)
		require.Nil(hm.add(2, startTime))
		require.Nil(hm.add(3, startTime.Local().AddDate(0, 0, 2)))
		require.Error(hm.add(4, startTime.Local().AddDate(0, 0, 1)))
		require.Nil(hm.add(5, startTime.Local().AddDate(0, 0, 3)))
	})
}
